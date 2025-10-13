
# Pergunta

NA AULA QUE EESTOU FAZENDO O CODIGO ABAIXO É UTILIZADO, SAO UTILIZADAS CHAVES PARA IMPRIMIR A ESTRUTURA DO STRUCT, O CURSO É ANTIGO, AINDA É ASSIM? O CURSO É DA ÉPOCA DE 2022, O QUE DEVO FICAR ATENTO QUE PODE SER MUDADO?

````go
package main

import (
    "fmt"
)

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	fmt.Println(ContaCorrente{})
}
````

## Resposta

Sim — esse código ainda é válido hoje. A “impressão com chaves” vem do **literal de struct** (`ContaCorrente{}`) e do formato padrão do `fmt`. Nada disso mudou desde 2022.

Aqui vai o que vale saber (e o que você pode atualizar):

### Como imprimir structs hoje

* `fmt.Println(ContaCorrente{})` usa o formato **%v** implícito → mostra só os **valores** dos campos, na ordem, entre chaves (sem nomes de campo). Ex.: `{ 0 0 0 0}`.
* Para ver **nomes dos campos**, use **`%+v`**:

  ```go
  fmt.Printf("%+v\n", ContaCorrente{})
  // {titular: numeroAgencia:0 numeroConta:0 saldo:0}
  ```
* Para ver a **representação Go-sintaxe** (bem útil em aula), use **`%#v`**:

  ```go
  fmt.Printf("%#v\n", ContaCorrente{})
  // main.ContaCorrente{titular:"", numeroAgencia:0, numeroConta:0, saldo:0}
  ```

### Campos exportados x não exportados (importante em 2022 e hoje)

Se depois você quiser (de)serializar com `encoding/json` ou acessar campos de outro pacote, eles **precisam ser exportados** (primeira letra maiúscula). Seu exemplo usa minúsculas (não exportados). Em JSON, por exemplo, esses campos **não saem**.

```go
type ContaCorrente struct {
    Titular       string  // exportado
    NumeroAgencia int     // exportado
    NumeroConta   int
    Saldo         float64
}
```

### O que mudou no ecossistema desde 2022 (e não quebra sua aula)

* **Generics (Go 1.18+)**: chegaram em 2022, mas **não afetam** structs simples. Ótimos para coleções/funções genéricas; você pode ignorar no primeiro passo.
* **Novidades recentes (1.21–1.23)**: coisas como `log/slog`, otimizações, e **`range` sobre inteiros** (`for i := range 10`) apareceram, mas **não impactam** a parte de structs/impressão.
* **`encoding/json`** continua igual no ponto central: **só vê campos exportados**. Dica: use tags `json:"nome_do_campo"` quando chegar lá.

### Boas práticas que valem para 2022 e agora

* Zero values continuam os mesmos: `""`, `0`, `0.0`, `false`, `nil` — então `ContaCorrente{}` está ok para iniciar “zerado”.
* Para **alterar saldo com métodos**, prefira **receiver por ponteiro**:

  ```go
  func (c *ContaCorrente) Depositar(v float64) { c.Saldo += v }
  ```
* Se quiser manter invariantes (ex.: não deixar saldo negativo), **não exporte** `Saldo` e ofereça métodos (`Sacar`, `Depositar`).

### Exemplo “atualizado” e didático

```go
package main

import "fmt"

type ContaCorrente struct {
    Titular       string
    NumeroAgencia int
    NumeroConta   int
    Saldo         float64
}

func (c *ContaCorrente) Depositar(v float64) { c.Saldo += v }

func main() {
    // zero value
    fmt.Printf("%#v\n", ContaCorrente{})

    // com valores
    cc := ContaCorrente{Titular: "Ana", NumeroAgencia: 123, NumeroConta: 456}
    cc.Depositar(100)
    fmt.Printf("%+v\n", cc) // mostra nomes dos campos
}
```

**Resumo:** o uso de chaves e a forma de imprimir structs com `fmt` seguem iguais. O que realmente pega para não “quebrar” é lembrar de **%+v / %#v** e **exportação de campos** quando for integrar com outros pacotes (como JSON).

