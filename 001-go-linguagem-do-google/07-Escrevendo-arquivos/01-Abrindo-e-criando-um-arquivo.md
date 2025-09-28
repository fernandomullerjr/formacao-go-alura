# 01 - Abrindo e criando um arquivo

---

## Transcrição

Com o monitoramento concluído, falta implementarmos a opção do nosso programa de escrever logs. Para isso, teremos que escrever em arquivos, para registrar em um arquivo de texto quando o site estava online e quando estava offline, inclusive o horário em que isso ocorrer.

Para isso, vamos criar a função `registraLog`:

```go
// restante do código omitido

func registraLog() {

}
```

Essa função tem que salvar o site e o seu status, logo ela precisa ter acesso a esses dois valores. Então, vamos recebê-los por parâmetro, sendo o status um booleano (true para o site online e false para offline):

```go
// restante do código omitido

func registraLog(site string, status bool) {

}
```

Agora, na função `testaSite`, se a o site for carregado com sucesso, nós chamamos a função `registraLog`, passando o valor true, mas se não for carregado com sucesso, passamos o valor false:

```go
// restante do código omitido

func testaSite(site string) {

    resp, err := http.Get(site)

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
        registraLog(site, true)
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
        registraLog(site, false)
    }
}
```

Agora, ao invés de criarmos o arquivo na mão, podemos pedir para o Go fazer isso, mas como?

---

## Criando um arquivo com Go

Para o Go criar um arquivo, podemos utilizar a função `OpenFile`, também do pacote `os`. Ela recebe o nome do arquivo, uma flag para representar o que fazer com o arquivo, e o seu tipo de permissão. Há diversas flags que podemos utilizar, elas podem ser vistar aqui.

No nosso caso, se o arquivo não existir, queremos criá-lo, então vamos utilizar a flag `O_CREATE` ou `O_RDWR`, para ler e escrever no arquivo. Além disso, vamos dar a permissão `0666` para ele:

```go
// restante do código omitido

func registraLog(site string, status bool) {

    arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR, 0666)

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    fmt.Println(arquivo)
}
```

Com isso, já conseguimos criar o arquivo, falta escrever nele. Veremos como fazer isso no próximo vídeo.





- Efetuando teste

>
> go run 01-abrindo-arquivo-e-criando-arquivo.go
Olá, sr. Douglas
Este programa está na versão 1.2
&{0xc00012c180}
1- Iniciar Monitoramento
2- Exibir Logs
0- Sair do Programa

- Trouxe o ponteiro ali, indicando o arquivo.



## RESUMO
Olá! Nesta aula, aprendemos como implementar a função registraLog em Go para registrar o status de um site em um arquivo de texto.

Primeiro, a função registraLog é criada para receber o site e seu status como parâmetros. Em seguida, dentro da função testaSite, a função registraLog é chamada com o status do site (true para online, false para offline).

Para criar o arquivo de log, a função OpenFile do pacote os é utilizada, especificando o nome do arquivo, as flags O_CREATE e O_RDWR para criar o arquivo se não existir e permitir leitura e escrita, e a permissão 0666.
    arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR, 0666)

No final da aula, o arquivo é criado, mas ainda falta aprender como escrever nele, o que será abordado no próximo vídeo.



