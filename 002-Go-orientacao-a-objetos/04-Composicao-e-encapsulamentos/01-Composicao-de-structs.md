# 01 Composição de structs

## Transcrição

Criamos a struct de Conta Corrente com alguns métodos como Sacar(), Depositar() e Transferir().

Mas existe um ponto importante para o qual ainda não demos muita atenção, o Titular. A string com o nome é a única informação que temos sobre a pessoa responsável por manter essa conta corrente. Precisamos de mais informações.

Poderíamos criar as variáveis TitularCPF, e TitularProfissao, ambas do tipo string. Mas seria um problema definir várias características de um titular em ContaCorrente. Nosso banco é digital, outros tipos de contas aparecerão e haverá outros titulares para outros tipos de contas.

Então, estaríamos passando a responsabilidade de manter um titular de uma conta para ContaCorrente, uma função que não deve ser dela. Pensando nisso, criaremos uma nova struct para manter os titulares de contas. Usaremos os mesmos campos, mas assim conseguiremos um código mais organizado e bem dividido de acordo com responsabilidades.

Criaremos um novo folder ou pasta chamada "clientes". No pacote "clientes", criaremos um novo arquivo, "cliente.go". Diremos que esse arquivo pertence ao package clientes e vamos declarar a nova struct com type Titular struct.

Já sabemos que Titular terá um Nome, CPF e Profissao do tipo string. Já deixamos a visibilidade dos campos acessível para outras partes do código.

```go
package clientes

type Titular struct {
    Nome      string
    CPF       string
    Profissao string
}
```

Agora vamos na contaCorrente e no lugar de dizer que Titular é uma string, diremos que ele é do tipo Titular. Porém, esse tipo virá de outro pacote. Por isso será necessária a importação. Digitaremos pwd no terminal para ganhar tempo encontrando o caminho para a pasta onde está o arquivo e copiar de "github" em diante.

```go
package contas

import "github.com/alura/banco/clientes"

type ContaCorrente struct {
    Titular       clientes.Titular
    NumeroAgencia int
    NumeroConta   int
    Saldo         float64
}
```

Vamos salvar e não haverá nenhuma mensagem de erro. Dessa forma, criamos uma struct de Titular em "cliente.go" e em "contaCorrente.go" teremos outro tipo com a struct Titular dentro dele, com Titular clientes.Titular. Colocamos um tipo sobre outro tipo, o que chamamos de composição na linguagem Go. Em vez de uma herança, temos uma composição de structs.

A seguir veremos mais sobre como fazer essa composição, já que agora estamos usando duas estruturas, a struct de clientes e a de conta corrente.


## RESUMO
Nesta aula, aprendemos sobre a importância de organizar o código dividindo responsabilidades. 
Em vez de colocar todas as informações do titular diretamente na struct ContaCorrente, criamos uma nova struct chamada Titular no pacote "clientes" para armazenar os dados do titular (Nome, CPF e Profissão).
Em seguida, na struct ContaCorrente, substituímos o campo Titular (que antes era uma string) por um campo do tipo clientes.Titular. 
Isso é o que chamamos de composição: colocar uma struct dentro de outra.
A composição permite que a struct ContaCorrente contenha informações sobre o titular sem ter que se preocupar em armazenar e gerenciar esses dados diretamente. Isso torna o código mais organizado, fácil de entender e de manter.



## Dia 22/11/2025


>
> go mod tidy
go: finding module for package github.com/fernandomullerjr/formacao-go-alura/002-Go-orientacao-a-objetos/04-Composicao-e-encapsulamentos/01-composicao-de-structs/clientes
go: downloading github.com/fernandomullerjr/formacao-go-alura v0.0.0-20251123004833-9acaaa6adfd3
go: github.com/fernandomullerjr/formacao-go-alura/002-Go-orientacao-a-objetos/03-Retornos-pacotes-e-visibilidade/go_oo-aula3/contas imports
        github.com/fernandomullerjr/formacao-go-alura/002-Go-orientacao-a-objetos/04-Composicao-e-encapsulamentos/01-composicao-de-structs/clientes: module github.com/fernandomullerjr/formacao-go-alura@latest found (v0.0.0-20251123004833-9acaaa6adfd3), but does not contain package github.com/fernandomullerjr/formacao-go-alura/002-Go-orientacao-a-objetos/04-Composicao-e-encapsulamentos/01-composicao-de-structs/clientes
> date
Sat Nov 22 22:36:30 -03 2025