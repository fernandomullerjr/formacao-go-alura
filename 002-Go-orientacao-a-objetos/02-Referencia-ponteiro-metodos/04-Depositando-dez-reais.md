
# 04 Depositando dez reais

Para criar um método capaz de depositar R$10,00 no saldo de uma conta, foi desenvolvido o seguinte código:

~~~~go
package main

import (
    "fmt"
)

type Conta struct {
    saldo float64
}

func (c *Conta) depositarDezReais() float64 {
    return c.saldo + 10
}

func main() {
    contaTeste := Conta{saldo: 10}

    contaTeste.depositarDezReais()
    contaTeste.depositarDezReais()

    fmt.Println(contaTeste)
}
~~~~


## Pergunta
Com base no trecho de código acima, marque as alternativas verdadeiras.
Selecione uma alternativa

    Ao executar o código, uma mensagem de erro informará que o saldo está incorreto.

    Ao executar o código acima, teremos o saldo de 20 como resultado.

    Ao executar o código acima, teremos o saldo de 30 como resultado.

    Ao executar o código acima, teremos o saldo de 10 como resultado.



## Resposta
Ao executar o código acima, teremos o saldo de 10 como resultado.

Certo! Como não atribuímos o código no saldo da conta( c.saldo = c.saldo + 10), nosso saldo continuará sendo 10. Neste link, você pode executar o código acima código.
<https://go.dev/play/p/aG6sxPAuk2t>





## Corrigindo código

O problema no código é que o método `depositarDezReais()` está apenas **retornando** o valor calculado, mas não está **modificando** o saldo da conta. Vou corrigir:

````go
package main

import (
	"fmt"
)

type Conta struct {
	saldo float64
}

func (c *Conta) depositarDezReais() {
	c.saldo += 10
}

func main() {
	contaTeste := Conta{saldo: 10}

	fmt.Println("Saldo inicial:", contaTeste.saldo)
	
	contaTeste.depositarDezReais()
	fmt.Println("Após primeiro depósito:", contaTeste.saldo)
	
	contaTeste.depositarDezReais()
	fmt.Println("Após segundo depósito:", contaTeste.saldo)

	fmt.Println("Conta final:", contaTeste)
}
````

**Principais correções:**

1. **Mudei o método** para modificar diretamente o saldo: `c.saldo += 10`
2. **Removi o return** pois não precisamos retornar nada, apenas modificar
3. **Adicionei prints** para mostrar o progresso dos depósitos

**Resultado esperado:**
- Saldo inicial: 10
- Após primeiro depósito: 20  
- Após segundo depósito: 30
- Conta final: {30}

O receptor por ponteiro (`*Conta`) permite que o método modifique a struct original, ao invés de trabalhar com uma cópia.


Ótimas perguntas! Vou explicar:

## 1. Por que foi removido o `float64`?

O `float64` foi removido porque mudamos a **função do método**:

**Antes (versão com problema):**
```go
func (c *Conta) depositarDezReais() float64 {
    return c.saldo + 10  // Apenas calculava e retornava
}
```

**Depois (versão corrigida):**
```go
func (c *Conta) depositarDezReais() {
    c.saldo += 10  // Modifica diretamente o saldo
}
```

- **Antes**: O método **retornava** um valor `float64` (o saldo + 10), mas não modificava a conta
- **Depois**: O método **modifica** diretamente o saldo da conta e não precisa retornar nada (`void`)

## 2. Usar `c.saldo = c.saldo + 10` também funciona?

**Sim, perfeitamente!** Ambas as formas são equivalentes:

```go
// Forma 1: Operador de atribuição composta
c.saldo += 10

// Forma 2: Atribuição explícita  
c.saldo = c.saldo + 10
```

### Exemplo prático:

````go
package main

import "fmt"

type Conta struct {
    saldo float64
}

// Versão com +=
func (c *Conta) depositarComOperadorComposto() {
    c.saldo += 10
}

// Versão com = ... +
func (c *Conta) depositarComAtribuicaoExplicita() {
    c.saldo = c.saldo + 10
}

func main() {
    conta1 := Conta{saldo: 100}
    conta2 := Conta{saldo: 100}
    
    conta1.depositarComOperadorComposto()
    conta2.depositarComAtribuicaoExplicita()
    
    fmt.Println("Conta1:", conta1.saldo) // 110
    fmt.Println("Conta2:", conta2.saldo) // 110
}
````

### Outros operadores compostos equivalentes:

```go
c.saldo += 10   // equivale a: c.saldo = c.saldo + 10
c.saldo -= 5    // equivale a: c.saldo = c.saldo - 5
c.saldo *= 2    // equivale a: c.saldo = c.saldo * 2
c.saldo /= 3    // equivale a: c.saldo = c.saldo / 3
```

**Resumo**: O `+=` é apenas uma forma mais concisa e idiomática de escrever a mesma operação!



- Testando o código:

> go run 04-depositando-dez-reais.go
{10}
> go run 04-versao-corrigida-depositando-dez-reais.go
Saldo inicial: 10
Após primeiro depósito: 20
Após segundo depósito: 30
Conta final: {30}
> date
Sun Nov  9 21:03:28 -03 2025
> pwd
/home/fernando/cursos/golang/formacao-go-alura/002-Go-orientacao-a-objetos/02-Referencia-ponteiro-metodos