# 05 - Lendo de um arquivo

---

## 📝 Transcrição

Se o erro que está ocorrendo é que o `sites.txt` não existe, então vamos criá-lo dentro da pasta `hello`, com o conteúdo que estava no nosso slice:

```
https://random-status-code.herokuapp.com/
https://www.alura.com.br
https://www.caelum.com.br
```

Agora, ao executar o programa e iniciar o monitoramento, vemos um código estranho sendo impresso. Isso acontece pois o que a função `Open` nos retorna nada mais é do que um ponteiro para o arquivo em si.

Então, precisamos ler o arquivo de outro jeito, e há diversas maneiras para isso. Uma delas, um jeito fácil e rápido para ler um arquivo inteiro, sem ter que ler linha a linha, é utilizar o pacote `os`.

---

### Lendo os dados de um arquivo

Através do pacote `os`, chamamos a função `ReadFile`, para ler o arquivo passado:

```go
func leSitesDoArquivo() []string {
    var sites []string
    arquivo, err := os.ReadFile("sites.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    fmt.Println(arquivo)

    return sites
}
```

Essa função nos retorna o arquivo em um array de bytes, então basta convertê-los para uma string através da função `string`:

```go
func leSitesDoArquivo() []string {
    var sites []string
    arquivo, err := ioutil.ReadFile("sites.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    fmt.Println(string(arquivo))

    return sites
}
```

Ao executar o programa e iniciar o monitoramento, vemos o conteúdo do arquivo sendo impresso no console! Mas esse é um jeito bom para quem quer exibir o conteúdo todo do arquivo, o que não é muito útil para nós, já que queremos ler cada site de uma vez, para salvar cada um dentro do slice.

---

### Lendo a primeira linha do arquivo

Para ler o arquivo linha a linha, vamos ver uma terceira forma de ler arquivos em Go, utilizando o pacote `bufio`. Para isso, vamos voltar à primeira leitura do arquivo, utilizando `os.Open`:

```go
func leSitesDoArquivo() []string {
    var sites []string
    arquivo, err := os.Open("sites.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    return sites
}
```

E com o `bufio`, nós criamos um leitor através da função `NewReader`, que recebe o arquivo a ser lido:

```go
func leSitesDoArquivo() []string {
    var sites []string
    arquivo, err := os.Open("sites.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)

    return sites
}
```

Com esse leitor, temos funções que nos ajudam a ler o arquivo, lendo linha a linha, como por exemplo a `ReadString`, que lê uma linha do arquivo e nos retorna em forma de string. Essa função deve receber um byte delimitador, para saber até onde queremos ler a linha.

No nosso caso, queremos ler a linha inteira, ou seja, até a quebra de linha, que é representada pelo `\n`. Como queremos representar um byte, utilizaremos aspas simples:

```go
func leSitesDoArquivo() []string {
    var sites []string
    arquivo, err := os.Open("sites.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)

    leitor.ReadString('\n')

    return sites
}
```

Essa função nos retorna a linha e um possível erro, que vamos detectar. Além disso, vamos imprimir a linha:

```go
func leSitesDoArquivo() []string {
    var sites []string

    arquivo, err := os.Open("sites.txt")
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    //Com esse leitor, temos funções que nos ajudam a ler o arquivo, lendo linha a linha, como por exemplo a `ReadString`, que lê uma linha do arquivo e nos retorna em forma de string. Essa função deve receber um byte delimitador, para saber até onde queremos ler a linha.
    //No nosso caso, queremos ler a linha inteira, ou seja, até a quebra de linha, que é representada pelo `\n`. Como queremos representar um byte, utilizaremos aspas simples:
    leitor := bufio.NewReader(arquivo)

    linha, err := leitor.ReadString('\n')
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    fmt.Println(linha)

    return sites
}
```

Quando iniciamos o monitoramento, a primeira linha do arquivo é impressa. Mas queremos ler todas as linhas do arquivo, então como fazemos isso? Veremos no próximo vídeo.





## Testando


> go run 01-arquivos-em-go.go
&{0xc00008e240}

ao executar o programa e iniciar o monitoramento, vemos um código estranho sendo impresso. Isso acontece pois o que a função `Open` nos retorna nada mais é do que um ponteiro para o arquivo em si.
quando o código exibe este código
é porque a os.Open traz o arquivo direto, não pega o seu conteúdo
este código é o arquivo no processador direto, ponteiro para o arquivo em si


## IMPORTANTE
- No video desta aula é mostrado o pacote ioutil
ele foi descontinuado/depreciado, conforme:
<https://pkg.go.dev/io/ioutil>


Sim! Desde o Go 1.16 o pacote `ioutil` foi **descontinuado** e suas funções migraram para `io` e `os`. Eis o mapa de substituições:

| `ioutil`           | Equivalente                                                        |
| ------------------ | ------------------------------------------------------------------ |
| `ioutil.ReadFile`  | `os.ReadFile`                                                      |
| `ioutil.WriteFile` | `os.WriteFile`                                                     |
| `ioutil.ReadAll`   | `io.ReadAll`                                                       |
| `ioutil.NopCloser` | `io.NopCloser`                                                     |
| `ioutil.Discard`   | `io.Discard`                                                       |
| `ioutil.ReadDir`   | `os.ReadDir` *(retorna `[]os.DirEntry` em vez de `[]os.FileInfo`)* |
| `ioutil.TempDir`   | `os.MkdirTemp`                                                     |
| `ioutil.TempFile`  | `os.CreateTemp`                                                    |

### Exemplos rápidos

**Ler arquivo**

```go
// antes
b, err := ioutil.ReadFile("arq.txt")
// depois
b, err := os.ReadFile("arq.txt")
```

**Escrever arquivo**

```go
// antes
err := ioutil.WriteFile("arq.txt", dados, 0644)
// depois
err := os.WriteFile("arq.txt", dados, 0644)
```

**Ler todo o corpo de um io.Reader (ex.: HTTP)**

```go
// antes
body, err := ioutil.ReadAll(resp.Body)
// depois
body, err := io.ReadAll(resp.Body)
```

**Listar diretório**

```go
// antes
fis, err := ioutil.ReadDir(".") // []os.FileInfo
// depois
entries, err := os.ReadDir(".") // []os.DirEntry
for _, e := range entries {
    info, _ := e.Info() // se precisar de os.FileInfo
    _ = info
}
```

**Arquivos/pastas temporários**

```go
// antes
f, err := ioutil.TempFile("", "prefix-")
// depois
f, err := os.CreateTemp("", "prefix-")

// antes
dir, err := ioutil.TempDir("", "prefix-")
// depois
dir, err := os.MkdirTemp("", "prefix-")
```

**Descartar escrita / criar ReadCloser neutro**

```go
w := io.Discard           // no lugar de ioutil.Discard
rc := io.NopCloser(rdr)   // no lugar de ioutil.NopCloser
```

Basta remover `import "io/ioutil"` e, conforme o caso, **importar `io` e/ou `os`**.



## IMPORTANTE
- No video desta aula é mostrado o pacote ioutil
ele foi descontinuado/depreciado, conforme:
<https://pkg.go.dev/io/ioutil>

- Suas funções migraram para `io` e `os`.