# 03 - Detectando erros

---

## 📝 Transcrição

Para tratar os erros, primeiramente devemos substituir o operador de identificador em branco (`_`) pela variável `err`, padrão para os erros em Go:

```go
func leSitesDoArquivo() []string {
    var sites []string
    arquivo, err := os.Open("sites.txt")
    fmt.Println(arquivo)
    return sites
}
```

E agora, se `err` não for nulo, significa que houve um erro, então vamos imprimi-lo:

```go
func leSitesDoArquivo() []string {
    var sites []string
    arquivo, err := os.Open("sites.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    fmt.Println(arquivo)
    return sites
}
```

Agora, ao executar o programa e iniciar o monitoramento, vemos a seguinte impressão:

```
Ocorreu um erro: open sites.txt: no such file or directory
```

Agora está bem claro o que aconteceu: o arquivo ou o diretório não foi encontrado.

---

Agora, vamos tratar do outro erro que estamos ignorando, no momento em que fazemos a requisição para o site, na função `testaSite`:

```go
func testaSite(site string) {
    resp, err := http.Get(site)

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}
```

Agora que já estamos detectando os erros do programa, vamos prosseguir com a leitura do arquivo.


## resumo
Olá! Tudo bem?Nesta aula, o foco principal é aprimorar o tratamento de erros em Go.Inicialmente, a variável _ (identificador em branco) é substituída por err ao abrir o arquivo "sites.txt". Em seguida, é implementada uma verificação para identificar se err é diferente de nil, indicando a ocorrência de um erro. Caso um erro seja detectado, ele é impresso no console, fornecendo informações claras sobre o problema, como a ausência do arquivo ou diretório.Além disso, a aula aborda a detecção de erros durante as requisições para os sites na função testaSite. Similarmente, é verificado se a requisição retorna um erro e, em caso afirmativo, o erro é impresso no console.O objetivo é garantir que o programa seja capaz de identificar e lidar adequadamente com os erros que possam surgir durante a leitura do arquivo e as requisições aos sites, fornecendo informações úteis para o usuário ou para fins de depuração.Ficou alguma dúvida?






## TESTANDO

>
> go run 03-detectando-erros.go
Ocorreu um erro: open sites.txt: no such file or directory
Estes são os sites que serão monitorados:

1- Iniciar Monitoramento
2- Exibir Logs
0- Sair do Programa
^Csignal: interrupt