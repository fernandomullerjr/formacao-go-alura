
# Capital do Estado

Manoela está fazendo um programa para imprimir informações sobre determinada cidade. Por enquanto o código está assim:

~~~~go
package main

import "fmt"

func main() {
    cidade, populacao, capital := devolveCidadeEPopulacao()
    if capital {
        fmt.Println("A capital ", cidade, "tem", populacao, "habitantes")
    } else {
        fmt.Println("A cidade ", cidade, "tem", populacao, "habitantes")
    }
}

func devolveCidadeEPopulacao() (string, int, bool) {
    return "Vila Sem Nome", 4328, true
}
~~~~

Quando ela executar o programa, o que será impresso no terminal?


R:
A capital Vila Sem Nome tem 4328 habitantes

Você acertou em cheio!

Fernando, sua resposta está correta, parabéns! O programa define uma função devolveCidadeEPopulacao que retorna o nome da cidade, a população e um valor booleano indicando se é capital. A função main chama essa função e imprime uma mensagem formatada com as informações da cidade, incluindo a palavra "capital" se o valor booleano for verdadeiro. Continue praticando e aprofundando seus conhecimentos em Go, você está no caminho certo!