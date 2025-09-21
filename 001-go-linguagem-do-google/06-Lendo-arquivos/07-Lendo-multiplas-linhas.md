# 07 - Lendo múltiplas linhas

---

## Transcrição

Se queremos que o nosso programa leia além da primeira linha do arquivo, o código responsável por fazer a leitura deve ser executado mais de uma vez. Então, o que vamos fazer é colocar esse código dentro de um `for`, até dar um erro específico, o erro de EOF (End Of File), que acontece quando não há mais linhas a serem lidas.

---

### Exemplo de leitura de múltiplas linhas

```go
func leSitesDoArquivo() []string {
    var sites []string

    arquivo, err := os.Open("sites.txt")
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)

    for {
        linha, err := leitor.ReadString('\n')
        fmt.Println(linha)
        if err == io.EOF {
            fmt.Println("Ocorreu um erro:", err)
        }
    }

    return sites
}
```

---

## Como sair do loop?

Se houver o erro, que já estamos verificando, nós damos um `break`, que faz com que o código saia do loop:

```go
func leSitesDoArquivo() []string {
    var sites []string

    arquivo, err := os.Open("sites.txt")
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)

    for {
        linha, err := leitor.ReadString('\n')
        fmt.Println(linha)
        if err == io.EOF {
            break
        }
    }

    return sites
}
```

---

Ao iniciar o monitoramento, vemos a seguinte impressão:

```
https://random-status-code.herokuapp.com/

https://www.alura.com.br

https://www.caelum.com.br
```

A impressão está sendo desse jeito pois estamos pulando linha no arquivo, então ela é lida também. O leitor lê a linha, incluindo o `\n`, por isso que fica esse pulo de linha na hora de impressão.

---

## Removendo quebras de linha

Devemos remover essa quebra de linha da linha lida, antes de adicioná-la ao slice. Para isso, existe a função `TrimSpace`, do pacote `strings`, que remove as quebras de linha e espaços ao final de uma string:

```go
func leSitesDoArquivo() []string {
    var sites []string

    arquivo, err := os.Open("sites.txt")
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)

    for {
        linha, err := leitor.ReadString('\n')
        linha = strings.TrimSpace(linha)
        fmt.Println(linha)
        if err == io.EOF {
            break
        }
    }

    return sites
}
```

---

## Adicionando ao slice

Agora, as linhas são impressas sem quebra de linha. Logo, já podemos adicioná-las ao slice:

```go
func leSitesDoArquivo() []string {
    var sites []string

    arquivo, err := os.Open("sites.txt")
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)

    for {
        linha, err := leitor.ReadString('\n')
        linha = strings.TrimSpace(linha)
        sites = append(sites, linha)
        if err == io.EOF {
            break
        }
    }

    return sites
}
```

---

Ao executar o programa novamente e iniciar o monitoramento, podemos perceber que os sites são monitorados corretamente! Agora, se quisermos adicionar mais sites, basta colocá-los no arquivo `sites.txt`.

---

## Fechando o arquivo

Por último, devemos ser educados com o sistema operacional. Se abrimos um arquivo com `os.Open`, após lê-lo, é uma boa prática fechá-lo com a função `Close`:

```go
func leSitesDoArquivo() []string {
    var sites []string

    arquivo, err := os.Open("sites.txt")
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)

    for {
        linha, err := leitor.ReadString('\n')
        linha = strings.TrimSpace(linha)
        sites = append(sites, linha)
        if err == io.EOF {
            break
        }
    }

    arquivo.Close()

    return sites
}
```

---


## RESUMO

Olá! Nesta aula, aprendemos como ler múltiplas linhas de um arquivo em Go. Para isso, utilizamos um loop for que lê o arquivo linha por linha até encontrar o erro io.EOF (End Of File), que indica o final do arquivo.
Além disso, vimos como remover a quebra de linha que o bufio.NewReader inclui ao ler cada linha, utilizando a função strings.TrimSpace. 
Também aprendemos a adicionar cada linha lida a um slice de strings e a fechar o arquivo após a leitura, utilizando a função arquivo.Close().
