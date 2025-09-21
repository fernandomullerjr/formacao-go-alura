# 06 - Obtendo o conteúdo de um arquivo

---

Observe o código abaixo, que abre o arquivo existente `"sites.txt"`:

```go
package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  arquivo, err := os.Open("sites.txt")
  if err != nil {
    fmt.Println("Ocorreu um erro:", err)
  }
  leitor := bufio.NewReader(arquivo)
}
```

---

## O que você deve fazer para ler a primeira linha do arquivo?

### Alternativas

#### ✅ Correta

```go
linha, _ := leitor.ReadString('\n')
```
**Isso aí!**

---

#### ❌ Alternativa incorreta

```go
linha, _ := bufio.ReadString(leitor, '\n')
```

---

#### ❌ Alternativa incorreta

```go
linha, _ := os.read(leitor)
```

---

#### ❌ Alternativa incorreta

```go
linha, _ := leitor.ReadString()
```

---

**Parabéns, você acertou!**