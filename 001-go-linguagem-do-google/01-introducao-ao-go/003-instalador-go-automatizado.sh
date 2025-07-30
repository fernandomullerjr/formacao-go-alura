#!/bin/bash

# ===============================
# Script de instalação do Go no Ubuntu (WSL) com configuração para ZSH
# Autor: Fernando Muller + ChatGPT
# ===============================

# ===== CONFIGURAÇÃO =====
GO_VERSION="1.24.4"
GO_TAR_FILE="go${GO_VERSION}.linux-amd64.tar.gz"
GO_DOWNLOAD_URL="https://go.dev/dl/${GO_TAR_FILE}"
INSTALL_DIR="/usr/local"
ZSHRC_FILE="$HOME/.zshrc"

# =========================

echo "🚀 Iniciando instalação do Go versão ${GO_VERSION}..."

# ✅ Removendo instalações anteriores
sudo rm -rf ${INSTALL_DIR}/go

# ✅ Baixando o Go
echo "⬇️  Baixando ${GO_TAR_FILE}..."
wget -c ${GO_DOWNLOAD_URL} -O ${GO_TAR_FILE}

if [ $? -ne 0 ]; then
    echo "❌ Erro no download do Go. Verifique sua conexão ou o link."
    exit 1
fi

# ✅ Instalando
echo "📦 Extraindo para ${INSTALL_DIR}..."
sudo tar -C ${INSTALL_DIR} -xzf ${GO_TAR_FILE}

# ✅ Configurando PATH no .zshrc
if grep -q "/usr/local/go/bin" ${ZSHRC_FILE}; then
    echo "✅ PATH do Go já configurado no ${ZSHRC_FILE}"
else
    echo "🔧 Adicionando Go ao PATH no ${ZSHRC_FILE}..."
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ${ZSHRC_FILE}
fi

# ✅ Aplicando as mudanças no shell atual
source ${ZSHRC_FILE}

# ✅ Verificando instalação
echo "🧠 Verificando instalação do Go..."
go version

if [ $? -eq 0 ]; then
    echo "🎉 Go instalado e configurado com sucesso!"
else
    echo "⚠️ Algo deu errado na instalação do Go."
fi

# ✅ Limpando o arquivo .tar.gz
rm -f ${GO_TAR_FILE}

echo "✅ Script finalizado."
