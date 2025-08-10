#!/usr/bin/env bash
# ============================================
# Instalação automatizada do Go (Linux geral)
# Testado em Debian 10/11/12, Ubuntu, etc.
# - Detecta versão estável mais recente (ou respeita $GO_VERSION)
# - Verifica SHA256 (com fallback caso o site entregue HTML)
# - Suporta amd64 e arm64
# - Configura PATH via /etc/profile.d (fallback para ~/.bashrc)
# Autor: Fernando Muller + ChatGPT
# ============================================

set -euo pipefail

INSTALL_DIR="/usr/local"
PROFILED_FILE="/etc/profile.d/go.sh"
USER_BASHRC="${HOME}/.bashrc"
TMP_DIR="$(mktemp -d)"
CLEANUP() { rm -rf "$TMP_DIR"; }
trap CLEANUP EXIT

command_exists() { command -v "$1" >/dev/null 2>&1; }
need_root() { [ "$(id -u)" -ne 0 ]; }

sudo_if_needed() {
  if need_root; then
    if command_exists sudo; then
      sudo "$@"
    else
      echo "❌ Precisa de root e 'sudo' não está disponível. Execute como root."
      exit 1
    fi
  else
    "$@"
  fi
}

fetch() {
  local url="$1" out="$2"
  if command_exists curl; then
    curl -fLSs "$url" -o "$out"
  elif command_exists wget; then
    wget -qO "$out" "$url"
  else
    echo "❌ Nem curl nem wget encontrados. Instale um deles e tente novamente."
    exit 1
  fi
}

detect_arch() {
  case "$(uname -m)" in
    x86_64|amd64) echo "amd64" ;;
    aarch64|arm64) echo "arm64" ;;
    armv7l|armv6l) echo "armv6l" ;;
    *) echo "❌ Arquitetura não suportada: $(uname -m)"; exit 1 ;;
  esac
}

is_hex_sha256() {
  [[ "$1" =~ ^[a-f0-9]{64}$ ]]
}

verify_sha256() {
  local tarball="$1" sha_file="$2"
  local expected computed
  expected="$(head -n1 "$sha_file" | awk '{print $1}' | tr -d ' \t\r\n')"

  # Se não vier um hash válido, avisa
  if ! is_hex_sha256 "$expected"; then
    echo "ℹ️  Checksum recebido não parece válido. Tentando rebaixar com ?download=1..."
    local url_dl="${SHA256_URL}?download=1"
    fetch "$url_dl" "$sha_file"
    expected="$(head -n1 "$sha_file" | awk '{print $1}' | tr -d ' \t\r\n')"
    if ! is_hex_sha256 "$expected"; then
      echo "⚠️  Ainda não consegui obter um SHA256 válido. Pulando verificação."
      return 0
    fi
  fi

  if command_exists sha256sum; then
    computed="$(sha256sum "$tarball" | awk '{print $1}')"
  elif command_exists shasum; then
    computed="$(shasum -a 256 "$tarball" | awk '{print $1}')"
  else
    echo "⚠️  sha256sum/shasum ausentes; pulando verificação."
    return 0
  fi

  if [ "$expected" != "$computed" ]; then
    echo "❌ SHA256 não confere!"
    echo "   Esperado: $expected"
    echo "   Obtido:   $computed"
    exit 1
  fi
  echo "✅ SHA256 conferido."
}

ARCH="$(detect_arch)"
if [ "${GO_VERSION:-}" = "" ]; then
  echo "🔍 Descobrindo última versão estável do Go..."
  if command_exists curl; then
    LATEST="$(curl -fsSL https://go.dev/VERSION?m=text | head -n1 || true)"
  else
    LATEST="$(wget -qO- https://go.dev/VERSION?m=text | head -n1 || true)"
  fi
  if [ -z "${LATEST}" ] || [[ ! "${LATEST}" =~ ^go[0-9] ]]; then
    echo "⚠️  Falha ao detectar versão; usando fallback go1.22.0."
    LATEST="go1.22.0"
  fi
  GO_VERSION="${LATEST#go}"
else
  GO_VERSION="${GO_VERSION#go}"
fi

GO_TAR_FILE="go${GO_VERSION}.linux-${ARCH}.tar.gz"
BASE_URL="https://go.dev/dl"
# Força ?download=1 para evitar páginas HTML
TARBALL_URL="${BASE_URL}/${GO_TAR_FILE}?download=1"
SHA256_URL="${BASE_URL}/${GO_TAR_FILE}.sha256"

TARBALL_PATH="${TMP_DIR}/${GO_TAR_FILE}"
SHA256_PATH="${TMP_DIR}/${GO_TAR_FILE}.sha256"

echo "🚀 Instalando Go ${GO_VERSION} (${ARCH})"

# Remover instalação anterior
if [ -d "${INSTALL_DIR}/go" ]; then
  echo "🧹 Removendo instalação anterior em ${INSTALL_DIR}/go..."
  sudo_if_needed rm -rf "${INSTALL_DIR}/go"
fi

# Download + checksum
echo "⬇️  Baixando ${GO_TAR_FILE}..."
fetch "${TARBALL_URL}" "${TARBALL_PATH}"

echo "🔐 Baixando checksum SHA256..."
fetch "${SHA256_URL}" "${SHA256_PATH}"

echo "🧪 Verificando SHA256..."
verify_sha256 "${TARBALL_PATH}" "${SHA256_PATH}"

# Instalação
echo "📦 Extraindo para ${INSTALL_DIR}..."
sudo_if_needed tar -C "${INSTALL_DIR}" -xzf "${TARBALL_PATH}"

# PATH
GO_PATH_LINE='export PATH=$PATH:/usr/local/go/bin'
APPLY_TO_SHELL=false

if sudo_if_needed test -d "$(dirname "${PROFILED_FILE}")"; then
  if [ ! -f "${PROFILED_FILE}" ] || ! sudo_if_needed grep -q "/usr/local/go/bin" "${PROFILED_FILE}"; then
    echo "🔧 Configurando PATH em ${PROFILED_FILE} (system-wide)..."
    printf "%s\n" "${GO_PATH_LINE}" | sudo_if_needed tee "${PROFILED_FILE}" >/dev/null
  else
    echo "✅ PATH do Go já presente em ${PROFILED_FILE}"
  fi
else
  if [ -f "${USER_BASHRC}" ] && grep -q "/usr/local/go/bin" "${USER_BASHRC}"; then
    echo "✅ PATH do Go já presente em ${USER_BASHRC}"
  else
    echo "🔧 Adicionando PATH do Go em ${USER_BASHRC}..."
    printf "\n# Go\n%s\n" "${GO_PATH_LINE}" >> "${USER_BASHRC}"
    APPLY_TO_SHELL=true
  fi
fi

# Aplicar no shell atual
export PATH="$PATH:/usr/local/go/bin"
if [ "${APPLY_TO_SHELL}" = true ] && [ -n "${BASH_VERSION:-}" ]; then
  # shellcheck disable=SC1090
  source "${USER_BASHRC}" || true
fi

# Teste
echo "🧠 Verificando instalação..."
if go version; then
  echo "🎉 Go instalado e configurado com sucesso!"
  echo "   Dica: abra um novo terminal para garantir o PATH system-wide."
else
  echo "⚠️ Algo deu errado na instalação do Go."
  exit 1
fi
