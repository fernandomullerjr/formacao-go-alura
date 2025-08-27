08
Percorrendo o slice
Próxima Atividade

Considere o código abaixo

package main

import "fmt"

func main() {
    pontosPlanningPoker := []int{1, 2, 3, 5, 8, 13, 21, 40}
    for i := 0; i < len(pontosPlanningPoker); i++ {
        fmt.Println(pontosPlanningPoker[i])
    }
}

Defina o conteúdo da próxima linha, de maneira que ela faça o programa compilar.

package main

import "fmt"

func main() {
    pontosPlanningPoker := []int{1, 2, 3, 5, 8, 13, 21, 40}
    for ________________________________ {
        fmt.Println("O ponto na posição", i, " tem o valor", ponto)
    }
}

Selecione uma alternativa

    i, ponto := range(pontosPlanningPoker)

    (ponto in pontosPlanningPoker)

    i, ponto : pontosPlanningP




## RESPOSTA

i, ponto := range(pontosPlanningPoker)

Isso aí!
