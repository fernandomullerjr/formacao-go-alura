# 01 Controle de fluxo com if

## Transcrição

Dando sequência à construção da nossa aplicação, agora que já conseguimos capturar o comando escolhido pelo usuário, podemos fazer algo de acordo com o comando digitado.

Se o usuário escolheu o comando 1, vamos fazer algo, se ele escolheu o comando 2, fazemos outra coisa, e assim por diante. Quando queremos saber se o usuário escolheu tal comando, precisamos utilizar a instrução if da programação.

A condição do if, no Go, não fica entre parênteses, como já é prática de outras linguagens:

~~~~go
if comando == 1 {

}
~~~~

E a condição deve sempre retornar um booleano, ou seja, true ou false. Como queremos testar os 3 comandos, vamos colocá-lo no if:

~~~~go
package main

import "fmt"

func main() {
    nome := "Douglas"
    versao := 1.1
    fmt.Println("Olá, sr.", nome)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    var comando int
    fmt.Scan(&comando)
    fmt.Println("O comando escolhido foi", comando)

    if comando == 1 {

    } else if comando == 2 {

    } else if comando == 0 {

    }
}
~~~~

E por último, se não for nenhum comando conhecido, vamos imprimir uma mensagem, demonstrando que não conhecemos o comando digitado. Vamos aproveitar e colocar as mensagens dos outros comandos:

~~~~go
package main

import "fmt"

func main() {
    nome := "Douglas"
    versao := 1.1
    fmt.Println("Olá, sr.", nome)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    var comando int
    fmt.Scan(&comando)
    fmt.Println("O comando escolhido foi", comando)

    if comando == 1 {
        fmt.Println("Monitorando...")
    } else if comando == 2 {
        fmt.Println("Exibindo Logs...")
    } else if comando == 0 {
        fmt.Println("Saindo do programa...")
    } else {
        fmt.Println("Não conheço este comando")
    }
}
~~~~

Com isso, concluímos os ifs, que não possui muito mistério, as diferenças para as outras linguagens é o não uso de parênteses na condição, que deve sempre retornar um booleano.



------------------------------------------------------------------------------
------------------------------------------------------------------------------
------------------------------------------------------------------------------
------------------------------------------------------------------------------
------------------------------------------------------------------------------
------------------------------------------------------------------------------
## RESUMO

Nesta aula, aprendemos como usar a instrução if em Go para controlar o fluxo do programa com base nas escolhas do usuário.

    if: Usado para executar diferentes blocos de código dependendo se uma condição é verdadeira (true) ou falsa (false).
    Condições sem parênteses: Em Go, as condições dentro do if não precisam estar entre parênteses.
    Múltiplas condições: Podemos usar else if para verificar múltiplas condições e else para um caso padrão quando nenhuma das condições anteriores for verdadeira.
