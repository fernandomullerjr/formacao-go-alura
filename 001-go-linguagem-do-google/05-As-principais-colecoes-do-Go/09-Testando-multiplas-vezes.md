# 09 Testando múltiplas vezes

## Transcrição

Queremos rodar o código que testa o site para cada um dos sites, então vamos exportá-lo ara uma função, a testaSite, que recebe o site por parâmetro:

~~~~go
// restante do código omitido

func testaSite(site string) {

    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}
~~~~

Agora chamamos essa função dentro do for, para cada item do nosso slice. Vamos aproveitar e diminuir a mensagem de teste e adicionar uma linha em branco, para dar um espaçamento entre as mensagens e o menu:

~~~~go
// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://random-status-code.herokuapp.com/", 
        "https://www.alura.com.br", "https://www.caelum.com.br"}

    for i, site := range sites {
        fmt.Println("Testando site", i, ":", site)
        testaSite(site)
    }

    fmt.Println("")
}
~~~~

Ao executar o programa, todos os sites do slice são monitorados.

Mas os sites são testados uma única vez e para testarmos novamente, devemos digitar o comando 1. Então, vamos fazer com que os sites sejam testados 5 vezes a cada monitoramento. Para isso, criamos mais um loop:

~~~~go
// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://random-status-code.herokuapp.com/",
        "https://www.alura.com.br", "https://www.caelum.com.br"}

    for i := 0; i < 5; i++ {
        for i, site := range sites {
            fmt.Println("Testando site", i, ":", site)
            testaSite(site)
        }        
    }

    fmt.Println("")
}
~~~~


## Aumentando o intervalo entre os monitoramentos

Mas o for executa muito rápido, então testaríamos cinco vezes cada site com um intervalo mínimo entre cada teste. Seria interessante termos um intervalo maior entre os testes, por exemplo de 5 em 5 minutos.

Para tal, a cada teste, podemos pedir para o Go esperar um pouco. Fazemos isso utilizando a função Sleep, do pacote time, passando para ela o quanto de tempo queremos esperar. Representamos o tempo através de constantes da própria biblioteca, como Second, Minute, entre outras:

~~~~go
// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://random-status-code.herokuapp.com/",
        "https://www.alura.com.br", "https://www.caelum.com.br"}

    for i := 0; i < 5; i++ {
        for i, site := range sites {
            fmt.Println("Testando site", i, ":", site)
            testaSite(site)
        }
        time.Sleep(5 * time.Minute)
    }
    fmt.Println("")
}
~~~~

No caso estamos testando de 5 em 5 minutos, mas esse tempo pode ser aumentado.


## Criando constantes

Por último, vamos nos livrar de alguns números do nosso código, e exportá-los para constantes. Por que constantes? Pois elas não podem ser modificadas.

Os números que queremos atacar é o número 5 que está dentro do for, que representa o número de monitoramentos, e o número 5 dentro do Sleep, que representa o delay do nosso monitoramento.

Então, após os imports, criamos as constantes:

~~~~go
package main

import (
    "fmt"
    "net/http"
    "os"
    "time"
)

const monitoramentos = 3
const delay = 5

// restante do código omitido
~~~~


Agora, na função iniciarMonitoramento, utilizamos essas constantes:

~~~~go
// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://random-status-code.herokuapp.com/",
        "https://www.alura.com.br", "https://www.caelum.com.br"}

    for i := 0; i < monitoramentos; i++ {
        for i, site := range sites {
            fmt.Println("Testando site", i, ":", site)
            testaSite(site)
        }

        time.Sleep(delay * time.Second)
    }

    fmt.Println("")
}
~~~~

Agora, se quisermos alterar o delay ou a quantidade de monitoramentos, basta alterarmos diretamente na declaração das constantes. E para finalizar, vamos adicionar mais alguns espaçamentos, a cada monitoramento dos nossos sites e depois que o comando foi escolhido:

~~~~go
// restante do código omitido

func leComando() int {
    var comandoLido int
    fmt.Scan(&comandoLido)
    fmt.Println("O comando escolhido foi", comandoLido)
    fmt.Println("")

    return comandoLido
}

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://random-status-code.herokuapp.com/",
        "https://www.alura.com.br", "https://www.caelum.com.br"}

    for i := 0; i < monitoramentos; i++ {
        for i, site := range sites {
            fmt.Println("Testando site", i, ":", site)
            testaSite(site)
        }
        time.Sleep(delay * time.Second)
        fmt.Println("")
    }
    fmt.Println("")
}
~~~~

Com isso, conseguimos monitorar mais de uma vez múltiplos sites.





- Testando:

~~~~bash
> go run 09-testando-multiplas-vezes-site.go
O meu slice tem 3 itens
O meu slice tem capacidade para 3 itens
O meu slice tem 4 itens
O meu slice tem capacidade para 6 itens
1
O comando escolhido foi 1
Monitorando...
Testando site 0 : https://random-status-code.herokuapp.com/
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x10 pc=0x625d65]

goroutine 1 [running]:
main.testaSite({0x714d21, 0x29})
        /home/fernando/cursos/golang/formacao-go-alura/001-go-linguagem-do-google/05-As-principais-colecoes-do-Go/09-testando-multiplas-vezes-site.go:107 +0x45
main.iniciarMonitoramento()
        /home/fernando/cursos/golang/formacao-go-alura/001-go-linguagem-do-google/05-As-principais-colecoes-do-Go/09-testando-multiplas-vezes-site.go:95 +0x27c
main.main()
        /home/fernando/cursos/golang/formacao-go-alura/001-go-linguagem-do-google/05-As-principais-colecoes-do-Go/09-testando-multiplas-vezes-site.go:22 +0xc5
exit status 2
>
> date
Wed Aug 27 19:11:41 -03 2025

 ~/cursos/g/formacao-go-alura/0/05-As-principais-colecoes-do-Go  main ?3    
~~~~


deu erro