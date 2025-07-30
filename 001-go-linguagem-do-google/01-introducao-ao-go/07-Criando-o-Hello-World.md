


O **passo a passo resumido** para criar e testar um projeto **Hello World em Go**, com build e execução:

---

## 🚀 **1. Criar projeto novo**

```bash
mkdir hello-go && cd hello-go
go mod init hello-go
```

> Isso cria o `go.mod` com o nome do módulo.

---

## ✍️ **2. Criar o código**

```bash
nano main.go
```

Cole o conteúdo:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

---

## ▶️ **3. Rodar direto (modo dev)**

```bash
go run main.go
```

✔️ Deve exibir:

```plaintext
Hello, Go!
```

---

## 🏗️ **4. Fazer o build (gerar binário)**

```bash
go build -o hello
```

Isso cria um binário chamado `hello`.

---

## ⚙️ **5. Executar o binário**

```bash
./hello
```

✔️ Saída:

```plaintext
Hello, Go!
```

---




- Criando o projeto do Hello World

Run 'go help mod init' for more information.
> go mod init hello-go
go: creating new go.mod: module hello-go
go: to add module requirements and sums:
        go mod tidy


- Executando o "go run" ele já compila e executa, sem precisar fazer o build:

> go run hello.go
Hello, World!


