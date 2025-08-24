# 01 Arrays em Go

## Transcrição

Neste capítulo, vamos evoluir o monitoramento do nosso programa, monitorando mais de um site ao mesmo tempo. Como programadores, não vamos criar uma variável para cada site que desejamos monitorar, e sim uma estrutura de dados responsável por conter várias strings, vários sites, o já conhecido Array.

Então, vamos ver como trabalhamos com coleções, principalmente Arrays e Slices, na linguagem Go.
Declarando um array

Primeiramente, vamos criar um array com a estrutura clássica, com o var. Para isso, além do var, nós devemos informar o nome do array, colchetes e o tipo de dados que ele guardará. Além disso, dentro dos colchetes, devemos informar o tamanho do array:

~~~~go
// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")
    var sites [4]string
    site := "https://www.alura.com.br"
    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}
~~~~

É importante nos atentar ao tipo de dados do array, não há espaço entre ele e os colchetes, na hora da declaração do array.
Colocando um valor dentro do array

Para colocar um valor no array, não há mistério, basta atribuir um valor a algum dos seus índices:

~~~~go
// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    var sites [4]string
    sites[0] = "https://random-status-code.herokuapp.com/"
    sites[1] = "https://www.alura.com.br"
    sites[2] = "https://www.caelum.com.br"

    site := "https://www.alura.com.br"
    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}
~~~~

Mas o fato do array ter um tamanho fixo, nos limita um pouco, pois se quisermos adicionar 5 itens no array, teremos que alterar o seu tamanho na sua declaração. Por isso, em Go, geralmente não trabalhamos com arrays, e sim com uma outra estrutura de dados, chamada Slice, que funciona em cima do array, mas não tem tamanho fixo.

Veremos mais sobre isso no próximo vídeo.



- Criando a versão completa:
formacao-go-alura/001-go-linguagem-do-google/05-As-principais-colecoes-do-Go/01-array-de-exemplo.go

- Testando:

~~~~bash
fernando@debian10x64:~/cursos/golang/formacao-go-alura/001-go-linguagem-do-google/05-As-principais-colecoes-do-Go$ go run 01-array-de-exemplo.go
Olá, sr. Douglas
Este programa está na versão 1.2
1- Iniciar Monitoramento
2- Exibir Logs
0- Sair do Programa
1
O comando escolhido foi 1
[https://random-status-code.herokuapp.com/ https://www.alura.com.br https://www.caelum.com.br ]
Monitorando...
Site: https://httpbin.org/status/404 está com problemas. Status Code: 404
1- Iniciar Monitoramento
2- Exibir Logs
0- Sair do Programa
0
O comando escolhido foi 0
Saindo do programa
fernando@debian10x64:~/cursos/golang/formacao-go-alura/001-go-linguagem-do-google/05-As-principais-colecoes-do-Go$
~~~~


## RESUMO

Nesta aula, aprendemos sobre Arrays em Go, que são coleções de dados de um mesmo tipo com tamanho fixo.

Você viu como declarar um array, especificando seu tamanho e tipo, e como atribuir valores a cada posição do array utilizando seus índices.

No entanto, a aula também destaca uma limitação dos arrays: seu tamanho fixo. Por isso, foi introduzida a estrutura de dados chamada Slice, que é mais flexível e utilizada com mais frequência em Go.