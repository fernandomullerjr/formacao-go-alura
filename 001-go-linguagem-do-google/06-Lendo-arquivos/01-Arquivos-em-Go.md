# 01 - Arquivos em Go

---

## 🎯 **Objetivo da Aula**

Monitorar sites a partir de um arquivo de texto, tornando o programa independente do slice criado manualmente.

---

## 📝 **Transcrição**

O que queremos fazer agora é não ficar dependente do slice, de só monitorar os sites contidos nele. Para isso, teremos um arquivo de texto com todos os sites que queremos monitorar. Adicionaremos vários sites em um arquivo de texto e ao iniciar nosso programa, ele monitorará todos os sites especificados dentro desse arquivo.

Para trabalhar com isso, devemos saber como trabalhamos com arquivo em Go, mais especificamente como lê-los. Para tal, criaremos a função `leSitesDoArquivo`, com essa responsabilidade. Além de ler os dados do arquivo, essa função os retornará em um slice:

```go
// restante do código omitido

func leSitesDoArquivo() []string {

}
```

Agora, na função `iniciarMonitoramento`, ao invés de criar o slice na mão, nós criamos a partir do retorno dessa função:

```go
// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    // criando o slice a partir da função leSitesDoArquivo()
    sites := leSitesDoArquivo()

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
```

Pronto, agora falta implementarmos a função. Antes de ler os dados do arquivo, devemos saber como abri-lo.

---

## 📂 **Abrindo um arquivo em Go**

Para abrir um arquivo em Go, precisamos pedir ao sistema operacional que abra o arquivo para nós. E quem é o representante do sistema operacional em Go? O pacote `os`, que já trabalhamos anteriormente. Nele, há a função `Open`, que abre o arquivo e nos retorna a sua representação em memória e um possível erro que possa ocorrer, que iremos ignorar:

```go
// restante do código omitido

func leSitesDoArquivo() []string {
    arquivo, _ := os.Open("sites.txt")
}
```

Vamos fazer um teste agora, vamos imprimir o arquivo, criar um slice vazio e retorná-lo:

```go
// restante do código omitido

func leSitesDoArquivo() []string {

    var sites []string
    arquivo, _ := os.Open("sites.txt")
    fmt.Println(arquivo)

    return sites
}
```

Agora, ao executar o programa e iniciar o monitoramento, vemos uma impressão `<nil>`. O que isso significa? Isso equivale ao `null` de outras linguagens de programação, ou seja, nosso arquivo está nulo. Isso acontece pois estamos abrindo um arquivo que não existe!

Mas como sabemos disso? Sabemos pois foi um erro proposital, já que estamos sendo descuidados, pois estamos ignorando os erros. Em nenhum momento estamos tratando os erros da linguagem de programação, já que, ao tentar abrir um arquivo que não existe, era para o sistema operacional nos avisar isso, mas estamos ignorando esses avisos, tanto no momento de abrir um arquivo, quanto no momento de realizar as requisições para os sites.

---

## ⏭️ **Próximo Passo**

Vamos ver como tratar esses erros no próximo vídeo.



## RESUMO
Nesta aula, aprendemos como ler dados de um arquivo de texto em Go para monitorar sites a partir de uma lista externa, em vez de um slice fixo no código.
1. Função leSitesDoArquivo: Criamos essa função para ler os sites de um arquivo e retorná-los em um slice.
2. Abrindo o arquivo: Usamos a função os.Open do pacote os para abrir o arquivo "sites.txt".
3. Tratamento de erros (breve): Foi mencionado que estamos ignorando os erros que podem ocorrer ao abrir o arquivo e ao fazer as requisições, o que não é uma boa prática. O tratamento de erros será abordado no próximo vídeo.
4. Teste inicial: Ao tentar abrir um arquivo que não existe, o programa retorna <nil>, indicando que o arquivo não foi encontrado.





> go run 01-arquivos-em-go.go
# command-line-arguments
./01-arquivos-em-go.go:94:24: undefined: sites
> date
Sat Sep 20 20:55:34 -03 2025

O erro ocorre porque a variável sites não está definida dentro da função iniciarMonitoramento().
Para corrigir, obtenha os sites usando a função leSitesDoArquivo() e utilize esse slice no loop.


> go run 01-arquivos-em-go.go
&{0xc00008e240}

O programa está apenas imprimindo a referência do arquivo aberto (fmt.Println(arquivo)), mas não está lendo os sites do arquivo sites.txt.
Para ler cada linha do arquivo e retornar um slice com os sites, atualize a função leSitesDoArquivo 