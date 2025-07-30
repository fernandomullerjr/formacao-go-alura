
# Importante - Mudanças

Curso da Alura indica uso de Workspace, que é o modelo antigo (usando $GOPATH) — ainda válido para projetos legados. 
O Go assume por padrão que seu workspace fica em:
```bash
$HOME/go
```

No contexto do **Go (Golang)** moderno — especialmente a partir do **Go 1.11+ com suporte a módulos (`go mod`)** — o conceito de **workspace** mudou em relação ao modelo antigo baseado no `GOPATH`.

---

# 📦 Onde fica o **workspace** do Go atualmente?

## ✅ **Com `go mod` (moderno e recomendado)**

Se você usa `go mod init` para seus projetos (e deve usar!), **não há mais um workspace fixo** como no modelo antigo. O seu workspace passa a ser **o diretório do seu projeto**, onde existe um arquivo `go.mod`.

### Exemplo:

```bash
~/cursos/go/formacao-go-alura/
├── go.mod
├── main.go
├── pkg/
└── internal/
```

> ✔️ Este diretório já é seu "workspace Go". Você pode ter vários workspaces, um por projeto.

---

## ⚠️ **Modelo antigo (usando `$GOPATH`) — ainda válido para projetos legados**

Se você **não usa `go mod`**, o Go assume por padrão que seu **workspace fica em:**

```bash
$HOME/go
```

Com a seguinte estrutura:

```bash
$GOPATH/
├── bin/       ← arquivos compilados executáveis
├── pkg/       ← arquivos .a de pacotes
└── src/       ← código-fonte Go (projetos ficam aqui)
```

Você pode verificar seu GOPATH com:

```bash
go env GOPATH
```

E seu workspace atual com:

```bash
go env GOWORK
```

> Mas se você usa módulos (`go mod`), **não precisa mais se preocupar com `GOPATH`** — apenas crie uma pasta, entre nela e use `go mod init`.

---

## 🧠 **Resumo prático:**

| Situação            | Caminho do Workspace                  |
| ------------------- | ------------------------------------- |
| Usando `go mod`     | O diretório do projeto com `go.mod`   |
| Não usando `go mod` | Valor de `$GOPATH`, geralmente `~/go` |

---

## 📌 Dicas para seu cenário (WSL + ZSH + Go 1.22+)

* Crie projetos onde quiser, por exemplo:

```bash
mkdir -p ~/projetos/go/meu-projeto
cd ~/projetos/go/meu-projeto
go mod init meu-projeto
```

* E tudo estará isolado no diretório local, sem depender de `GOPATH`.


