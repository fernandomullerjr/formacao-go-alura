package trabalhandocomvariaveis


Veja o código abaixo:

package main

import "fmt"

func main() {
    var nome string = "José"
    var peso float32 = 75.4
    fmt.Println("Olá, eu sou o", nome)
}

Esse programa gera um erro. Você sabe identificar a causa?



variável peso não é utilizada


Você acertou em cheio!

Fernando, sua resposta está correta, parabéns! O programa em Go declara a variável peso, mas ela não é utilizada em nenhuma parte do código. Em Go, declarar uma variável e não usá-la resulta em erro de compilação. Continue praticando e aprofundando seus conhecimentos em Go, você está no caminho certo!