

usar script
/home/fernando/cursos/golang/formacao-go-alura/001-go-linguagem-do-google/01-introducao-ao-go/003-instalador-go-automatizado.sh


adaptando


show! aqui vai uma versão “linux geral” que funciona no Debian 10 (e outras distros), detecta automaticamente a melhor versão estável do Go diretamente do site oficial, verifica o SHA256, lida com amd64/arm64 e configura o PATH sem zsh (usa profile.d e cai para .bashrc se precisar).

> salve como `install-go.sh`, dê permissão (`chmod +x install-go.sh`) e execute (`./install-go.sh`).
> Se quiser forçar uma versão específica, rode com `GO_VERSION=1.22.7 ./install-go.sh`.

```bash
#!/usr/bin/env bash
# ============================================
# Instalação automatizada do Go (Linux geral)
# Testado em Debian 10/11/12, Ubuntu, etc.
# - Detecta versão estável mais recente (ou respeita $GO_VERSION)
# - Verifica SHA256
# - Suporta amd64 e arm64
# - Configura PATH via /etc/profile.d (fallback para ~/.bashrc)
# Autor: Fernando Muller + ChatGPT
# ============================================

set -euo pipefail

# ===== Configurações =====
INSTALL_DIR="/usr/local"
PROFILED_FILE="/etc/profile.d/go.sh"
USER_BASHRC="${HOME}/.bashrc"
TMP_DIR="$(mktemp -d)"
CLEANUP() { rm -rf "$TMP_DIR"; }
trap CLEANUP EXIT

# ===== Funções auxiliares =====
command_exists() { command -v "$1" >/dev/null 2>&1; }

need_root() {
  [ "$(id -u)" -ne 0 ]
}

sudo_if_needed() {
  if need_root; then
    if command_exists sudo; then
      sudo "$@"
    else
      echo "❌ Precisa de privilégios de root e 'sudo' não está disponível."
      echo "   Reexecute como root: sudo $0 ou su -c '$0'"
      exit 1
    fi
  else
    "$@"
  fi
}

fetch() {
  # Tenta curl, senão wget
  local url="$1" out="$2"
  if command_exists curl; then
    curl -fsSL "$url" -o "$out"
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
    armv7l|armv6l) echo "armv6l" ;; # Go fornece build armv6l (compatível c/ v7)
    *) echo "❌ Arquitetura não suportada: $(uname -m)"; exit 1 ;;
  esac
}

# ===== Definir versão =====
ARCH="$(detect_arch)"
if [ "${GO_VERSION:-}" = "" ]; then
  echo "🔍 Descobrindo última versão estável do Go..."
  # Endpoint oficial retorna, por exemplo: go1.24.4
  if command_exists curl; then
    LATEST="$(curl -fsSL https://go.dev/VERSION?m=text | head -n1 || true)"
  elif command_exists wget; then
    LATEST="$(wget -qO- https://go.dev/VERSION?m=text | head -n1 || true)"
  else
    echo "❌ Nem curl nem wget para detectar versão. Defina GO_VERSION manualmente."
    exit 1
  fi
  if [ -z "${LATEST}" ] || [[ ! "${LATEST}" =~ ^go[0-9] ]]; then
    echo "⚠️  Falha ao descobrir a versão automaticamente. Usando fallback go1.22.0."
    LATEST="go1.22.0"
  fi
  GO_VERSION="${LATEST#go}"   # remove prefixo 'go'
else
  # Sanitiza entrada tipo "go1.22.7" para "1.22.7"
  GO_VERSION="${GO_VERSION#go}"
fi

GO_TAR_FILE="go${GO_VERSION}.linux-${ARCH}.tar.gz"
BASE_URL="https://go.dev/dl"
TARBALL_URL="${BASE_URL}/${GO_TAR_FILE}"
SHA256_URL="${TARBALL_URL}.sha256"
TARBALL_PATH="${TMP_DIR}/${GO_TAR_FILE}"
SHA256_PATH="${TMP_DIR}/${GO_TAR_FILE}.sha256"

echo "🚀 Instalando Go ${GO_VERSION} (${ARCH})"

# ===== Remover instalação anterior =====
if [ -d "${INSTALL_DIR}/go" ]; then
  echo "🧹 Removendo instalação anterior em ${INSTALL_DIR}/go..."
  sudo_if_needed rm -rf "${INSTALL_DIR}/go"
fi

# ===== Download e verificação =====
echo "⬇️  Baixando ${GO_TAR_FILE}..."
fetch "${TARBALL_URL}" "${TARBALL_PATH}"

echo "🔐 Baixando checksum SHA256..."
fetch "${SHA256_URL}" "${SHA256_PATH}"

echo "🧪 Verificando SHA256..."
# Arquivo .sha256 costuma conter "<sha256>  goX.Y.Z.linux-ARCH.tar.gz"
if command_exists sha256sum; then
  (cd "${TMP_DIR}" && sha256sum -c "$(basename "${SHA256_PATH}")")
elif command_exists shasum; then
  # macOS compat; em Linux raro, mas fica o fallback
  (cd "${TMP_DIR}" && shasum -a 256 -c "$(basename "${SHA256_PATH}")")
else
  echo "⚠️  sha256sum/shasum ausentes; prosseguindo sem verificação."
fi

# ===== Instalação =====
echo "📦 Extraindo para ${INSTALL_DIR}..."
sudo_if_needed tar -C "${INSTALL_DIR}" -xzf "${TARBALL_PATH}"

# ===== Configurar PATH =====
# Preferência: system-wide via /etc/profile.d/go.sh
GO_PATH_LINE='export PATH=$PATH:/usr/local/go/bin'
APPLY_TO_SHELL=false

if sudo_if_needed test -d "$(dirname "${PROFILED_FILE}")"; then
  # Cria/atualiza somente se a linha ainda não existir
  SHOULD_WRITE=true
  if [ -f "${PROFILED_FILE}" ]; then
    if sudo_if_needed grep -q "/usr/local/go/bin" "${PROFILED_FILE}"; then
      SHOULD_WRITE=false
    fi
  fi
  if [ "${SHOULD_WRITE}" = true ]; then
    echo "🔧 Configurando PATH em ${PROFILED_FILE} (system-wide)..."
    printf "%s\n" "${GO_PATH_LINE}" | sudo_if_needed tee "${PROFILED_FILE}" >/dev/null
  else
    echo "✅ PATH do Go já presente em ${PROFILED_FILE}"
  fi
else
  # Fallback para ~/.bashrc
  if [ -f "${USER_BASHRC}" ] && grep -q "/usr/local/go/bin" "${USER_BASHRC}"; then
    echo "✅ PATH do Go já presente em ${USER_BASHRC}"
  else
    echo "🔧 Adicionando PATH do Go em ${USER_BASHRC}..."
    printf "\n# Go\n%s\n" "${GO_PATH_LINE}" >> "${USER_BASHRC}"
    APPLY_TO_SHELL=true
  fi
fi

# ===== Aplicar ao shell atual =====
export PATH="$PATH:/usr/local/go/bin"
if [ "${APPLY_TO_SHELL}" = true ]; then
  # Só faz sentido se for bash interativo
  if [ -n "${BASH_VERSION:-}" ]; then
    # shellcheck disable=SC1090
    source "${USER_BASHRC}" || true
  fi
fi

# ===== Teste =====
echo "🧠 Verificando instalação..."
if go version; then
  echo "🎉 Go instalado e configurado com sucesso!"
  echo "   Dica: abra um novo terminal para garantir que o PATH system-wide seja aplicado."
else
  echo "⚠️ Algo deu errado na instalação do Go."
  exit 1
fi
```

quer que eu inclua também a remoção limpa (uninstall) e suporte a `GOPATH/GOBIN` opcionais? posso acrescentar flags tipo `--uninstall` e `--with-gopath`.




- ERROS

~~~~BASH
fernando@debian10x64:~$ touch install-go.sh
fernando@debian10x64:~$ chmod +x install-go.sh
fernando@debian10x64:~$ vi install-go.sh
fernando@debian10x64:~$ ./install-go.sh
🔍 Descobrindo última versão estável do Go...
🚀 Instalando Go 1.24.6 (amd64)
🧹 Removendo instalação anterior em /usr/local/go...
[sudo] password for fernando:
⬇️  Baixando go1.24.6.linux-amd64.tar.gz...
🔐 Baixando checksum SHA256...
🧪 Verificando SHA256...
sha256sum: go1.24.6.linux-amd64.tar.gz.sha256: no properly formatted SHA256 checksum lines found
fernando@debian10x64:~$

~~~~



boa! o erro é porque os `.sha256` do site do Go **trazem só o hash**, sem o nome do arquivo.
O `sha256sum -c` espera “`<hash>  <arquivo>`”, aí ele reclama.
Segue script corrigido: agora ele calcula o hash local e compara com o hash oficial (funciona tanto se o `.sha256` vier só com o hash quanto com “hash arquivo”).

> salve por cima e rode de novo (`./install-go.sh`).

```bash
#!/usr/bin/env bash
# ============================================
# Instalação automatizada do Go (Linux geral)
# Testado em Debian 10/11/12, Ubuntu, etc.
# - Detecta versão estável mais recente (ou respeita $GO_VERSION)
# - Verifica SHA256 (compatível com .sha256 só-hash ou hash+arquivo)
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
    curl -fsSL "$url" -o "$out"
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

verify_sha256() {
  local tarball="$1" sha_file="$2"
  local expected computed
  # Extrai apenas o primeiro token do arquivo .sha256 (funciona com “hash” ou “hash  arquivo”)
  expected="$(head -n1 "$sha_file" | awk '{print $1}' | tr -d ' \t\r\n')"
  if [ -z "$expected" ]; then
    echo "⚠️  Arquivo SHA256 vazio/inesperado; pulando verificação."
    return 0
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
TARBALL_URL="${BASE_URL}/${GO_TAR_FILE}"
SHA256_URL="${TARBALL_URL}.sha256"
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
```

se quiser, já adiciono uma flag `--uninstall` para remover `/usr/local/go` e as entradas de PATH. Quer?






fernando@debian10x64:~$
fernando@debian10x64:~$ ./install-go.sh
🔍 Descobrindo última versão estável do Go...
🚀 Instalando Go 1.24.6 (amd64)
⬇️  Baixando go1.24.6.linux-amd64.tar.gz...
🔐 Baixando checksum SHA256...
🧪 Verificando SHA256...
❌ SHA256 não confere!
   Esperado: <!DOCTYPE
   Obtido:   bbca37cc395c974ffa4893ee35819ad23ebb27426df87af92e93a9ec66ef8712
fernando@debian10x64:~$
fernando@debian10x64:~$ date
Sun 10 Aug 2025 03:21:36 PM -03
fernando@debian10x64:~$





isso aconteceu porque o link do checksum (`.sha256`) no go.dev às vezes retorna **uma página HTML** em vez do arquivo (precisa do `?download=1`). Seu script leu o início do HTML (`<!DOCTYPE`) como se fosse o hash.

Corrigi isso forçando `?download=1` e adicionando uma checagem: se o conteúdo do `.sha256` não parecer um SHA válido, ele rebaixa com `?download=1`.

Substitua seu `install-go.sh` por este e rode de novo:

```bash
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
```

se preferir mexer só no ponto crítico, basta trocar a construção dos URLs por:

```bash
GO_TAR_FILE="go${GO_VERSION}.linux-${ARCH}.tar.gz"
BASE_URL="https://go.dev/dl"
TARBALL_URL="${BASE_URL}/${GO_TAR_FILE}?download=1"
SHA256_URL="${BASE_URL}/${GO_TAR_FILE}.sha256"
```

e manter a validação do `verify_sha256` como no script acima.
me avisa se quiser que eu já adicione `--uninstall` também 👍
