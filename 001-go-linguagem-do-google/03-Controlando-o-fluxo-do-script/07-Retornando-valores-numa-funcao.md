# 07 Retornando valores numa função

Observe o código abaixo.

~~~~go
package main

import "fmt"

func main() {
    palavraDigitada := lePalavra()
    fmt.Println("A palavra digitada foi", palavraDigitada)
}

func lePalavra() string {
    var palavra string
    fmt.Print("Digite uma palavra: ")
    fmt.Scan(&palavra)
    return palavra
}
~~~~

O que está faltando para que ele compile e a função retorne a palavra digitada para a função principal?
Selecione uma alternativa

    Nada, pois o código está compilando normalmente.

    Colocar o símbolo & na variável de retorno da função, para indicar que queremos passar a referência da mesma.

    return &palavra

    Colocar o tipo antes da declaração do nome da função.

    func string lePalavra() {


## RESPOSTA

Nada, pois o código está compilando normalmente.

Isso aí! A função lePalavra retorna uma string com o valor digitado pelo usuário.
