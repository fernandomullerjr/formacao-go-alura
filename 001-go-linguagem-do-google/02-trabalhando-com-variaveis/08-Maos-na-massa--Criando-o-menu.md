
# 08  Mãos na Massa: Criando o Menu
Próxima Atividade

    Começando deste ponto? Você pode fazer o DOWNLOAD completo do projeto do capítulo anterior e continuar seus estudos a partir deste capítulo.

Neste capítulo vamos criar a mensagem de introdução do script e o menu que será exibido para o usuário, além de capturar a escolha que ele fez.

1 - Para começar a trabalhar com variáveis, vamos declarar o nome do usuário e sua versão em duas variáveis, utilizando o operador declaração de variáveis curto:

//hello.go
package main

import "fmt"

func main() {
    nome := "Douglas"
    versao := 1.1
}

2 - Vamos exibir as mensagens de boas vindas com estas duas variáveis:

//hello.go
package main

import "fmt"

func main() {
    nome := "Douglas"
    versao := 1.1

    fmt.Println("Olá, sr(a).", nome)
    fmt.Println("Este programa está na versão", versao)
}

3 - Agora vamos exibir o Menu para o usuário escolher qual opção ele deseja. Serão 3 Println para exibir o Menu, um para cada opção:

//hello.go
package main

import "fmt"

func main() {
    nome := "Douglas"
    versao := 1.1

    fmt.Println("Olá, sr(a).", nome)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")
}

4- Com o menu sendo exibido corretamente, precisamos capturar a escolha do usuário do nosso programa. Vamos colocar a opção escolhida na variável int comando e capturar o seu valor com a função fmt.Scan():

//hello.go
package main

import "fmt"

func main() {
    nome := "Douglas"
    versao := 1.1

    fmt.Println("Olá, sr(a).", nome)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    var comando int
    fmt.Scan(&comando)
}

Atenção para não esquecer de passar o endereço (&) da variável comando para a função Scan

5- Por último imprima o valor da variável comando para garantir que estamos capturando com sucesso:

//hello.go
package main

import "fmt"

func main() {
    nome := "Douglas"
    versao := 1.1

    fmt.Println("Olá, sr(a).", nome)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    var comando int
    fmt.Scan(&comando)

    fmt.Println("O valor da variável comando é:", comando)
}

Conseguimos criar nossa mensagem de introdução e o menu de opções do usuário, além de exibir qual opção ele escolheu. Já estamos avançando em nosso programa !