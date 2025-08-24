# 03 Para saber mais: funções que retornam arrays
 
Considere o código abaixo.

~~~~go
package main

import (
    "fmt"
)

func main() {
    estados := devolveEstadosDoSudeste()
    fmt.Println(estados)
}

func devolveEstadosDoSudeste() [4]string {
    var estados [4]string
    estados[0] = "RJ"
    estados[1] = "SP"
    estados[2] = "MG"
    estados[3] = "ES"
    return estados
}
~~~~

O que será impresso no terminal?
Selecione uma alternativa

    Nada, porque o programa não compila, uma vez que não é possível retornar um array através de uma função.

    Nada, porque o programa não compila, uma vez que a declaração de retorno da função não exige a quantidade de posições do array.

    [RJ SP MG ES]


## RESPOSTA
[RJ SP MG ES]

Isso aí. O código acima cria uma função que retorna um array com quatro posições, que é atribuído a uma variável estado, e em seguida ela é impressa no terminal.
