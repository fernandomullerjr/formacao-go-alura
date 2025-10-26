# 05 Reflexão sobre structs

Para criar uma nova struct, usamos o prefixo type, seguido do nome da struct, o sufixo struct e declaramos os campos entre chaves, conforme o exemplo abaixo:

~~~~go
type Pessoa struct {
    nome string
    idade int
}
~~~~

Sabendo disso, analise as afirmações abaixo e marque as verdadeiras em relação ao uso de structs no Go.
Selecione 2 alternativas

    As variáveis ficam mais organizadas.

    Permite reutilizar o código.

    Permite criar uma nova pessoa com o código Pessoa(nome: "Luísa", idade: 28)

    Só permite criar um novo tipo informando o nome dos campos.


- Permite reutilizar o código.
Certo! Não precisamos criar 2 variáveis para criar um nova pessoa.

- Em Go, não utilizamos parênteses ( ) e sim chaves { }.
então a resposta3 não é correta