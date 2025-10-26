# 🧭 Introdução a Ponteiros em Go

Em Go, **ponteiros** são variáveis que **armazenam o endereço de memória** de outra variável.
Eles permitem **manipular valores diretamente**, sem precisar copiar dados.

---

## 📦 Conceito Básico

```go
var x int = 10
var p *int // ponteiro para int

p = &x     // & retorna o endereço de x
fmt.Println(p)  // exibe algo como 0xc0000140a8
fmt.Println(*p) // exibe 10 — desreferencia o ponteiro
```

* `&x` → obtém o **endereço** de `x`
* `*p` → acessa o **valor armazenado** nesse endereço

---

## 🔁 Modificando valores via ponteiro

```go
var nome string = "Fernando"
p := &nome

fmt.Println(*p) // "Fernando"

*p = "Muller"   // altera o valor no endereço
fmt.Println(nome) // "Muller"
```

> Alterar `*p` muda `nome`, porque ambos apontam para o mesmo local na memória.

---

## ⚙️ Ponteiros em funções

Sem ponteiro, o Go **copia o valor** (passagem por valor).
Com ponteiro, você **altera o original**.

```go
func dobrar(n *int) {
    *n = *n * 2
}

func main() {
    numero := 5
    dobrar(&numero)
    fmt.Println(numero) // 10
}
```

---

## 🚫 Nil e segurança

Todo ponteiro em Go **tem valor zero** igual a `nil` (nulo).

```go
var p *int
fmt.Println(p == nil) // true

// cuidado: acessar *p gera panic
// fmt.Println(*p) -> panic: invalid memory address
```

Sempre inicialize ponteiros antes de usar:

```go
x := 10
p := &x
```

---

## 🔧 Casos comuns de uso

* **Alterar variáveis dentro de funções** sem retorno.
* **Otimizar performance** (evitar cópias grandes).
* **Trabalhar com structs** e métodos (por exemplo, receptores com `*`).

---

## 🧱 Exemplo prático com struct

```go
type Conta struct {
    saldo float64
}

func (c *Conta) Depositar(valor float64) {
    c.saldo += valor
}

func main() {
    conta := Conta{saldo: 100.0}
    conta.Depositar(50.0)
    fmt.Println(conta.saldo) // 150.0
}
```

---

## 🧩 Resumo rápido

| Símbolo | Significado                          |
| ------- | ------------------------------------ |
| `&`     | endereço de uma variável             |
| `*`     | valor armazenado no endereço         |
| `nil`   | ponteiro nulo (não aponta para nada) |


