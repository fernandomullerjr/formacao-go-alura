# 04 - Capturando erros

---

Suponha que o arquivo `"sites.txt"` não existe. Quando o código abaixo é executado, a variável `err` conterá o erro de arquivo inexistente:

```go
arquivo, err := os.Open("sites.txt")
```

---

## O que o programador deve fazer para tratar esse erro?

### Alternativas

#### ✅ Correta

Verificar se a variável `err` é diferente de `nil`, e mostrar a mensagem de erro no terminal:

```go
if err != nil {
  fmt.Println("Ocorreu um erro:", err)
}
```

**Isso aí!**

---

#### ❌ Alternativa incorreta

Usar a estrutura `try capture`:

```go
try {
  arquivo, err := os.Open("sites.txt")
} capture {
  fmt.Println("Ocorreu um erro:", err)
}
```

---

#### ❌ Alternativa incorreta

Usar um `try catch`:

```go
try {
  arquivo, err := os.Open("sites.txt")
} catch {
  fmt.Println("Ocorreu um erro:", err)
}
```

---

> **Observação:**  
Go não possui estruturas `try/catch` ou `try/capture` para tratamento de erros. O padrão é sempre verificar se o erro retornado é diferente de `nil`.