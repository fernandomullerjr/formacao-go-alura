
# Cardápio apetitoso

Mário foi contratado por um restaurante para desenvolver um programa que solicita a escolha do cardápio. Ele implementou na linguagem Go e seu código está listado abaixo.

~~~~go
package main

import "fmt"

func main() {
    var prato string
    fmt.Println("Digite seu prato preferido...")
    fmt.Println("P - Pizza")
    fmt.Println("H - Hambúrguer")
    fmt.Println("B - Bife com fritas")
    fmt.Println("S - Salada Caesar")
    fmt.Println("F - Salada de Frutas")
    fmt.Println("E - Estrogonofe")
    fmt.Println("O - Outros")
    fmt.Scan(&prato)

    switch prato {
    case "B":
        fmt.Println("Com batatas Palito ou Noisete?")
    case "H":
        fmt.Println("Com Queijo ou com Ovo?")
    case "P":
        fmt.Println("Calabresa ou Napolitana?")
    case "S":
        fmt.Println("Alface ou Rúcula?")
    case "F":
        fmt.Println("Kiwi ou Frango?")
    case "E":
        fmt.Println("Carne ou Frango?")
    case "O":
        fmt.Println("Não gostou de nosso cardápio?")
    default:
        fmt.Println("Não entendi seu paladar...")
    }
}
~~~~

Quando o cliente digitar b, o que será impresso?


## RESPOSTA
Não entendi seu paladar...


## validando
Você acertou em cheio!

Fernando, sua resposta está correta, parabéns! O programa em Go compara a entrada do usuário com as opções de cardápio usando um switch. Como o usuário digitou "b" (minúsculo), e as opções no switch são em letras maiúsculas ("B", "H", "P", etc.), nenhuma das condições é satisfeita, e o programa executa o caso default, imprimindo "Não entendi seu paladar...". Continue praticando e explorando a linguagem Go, você está no caminho certo!