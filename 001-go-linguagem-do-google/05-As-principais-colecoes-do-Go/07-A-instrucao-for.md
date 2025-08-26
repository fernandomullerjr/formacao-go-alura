# 07 - A instrução for

## Transcrição

Vamos apagar a função `exibeNomes`, que foi útil para aprendermos o funcionamento do slice, e começar a utilizá-lo na nossa aplicação, para melhorar o armazenamento dos sites a serem monitorados:

```go
func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{
        "https://random-status-code.herokuapp.com/", 
        "https://www.alura.com.br", 
        "https://www.caelum.com.br",
    }

    site := "https://www.alura.com.br"
    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}
```

Agora que temos uma coleção de sites, podemos monitorar mais de um site ao mesmo tempo. Para isso, devemos percorrer os elementos do nosso slice, acessando e monitorando cada site.

---

### Percorrendo os itens do slice

Como vimos na aula anterior, não existe no Go outra estrutura de repetição além do `for`, então vamos utilizá-lo para percorrer os itens do slice de sites. Para cada item do slice, nós vamos mandar uma requisição e testar o seu status code.

Vimos o `for` infinito, que faz com que o código seja repetido para sempre, mas também podemos utilizar o `for` "tradicional", onde declaramos uma variável e vamos incrementando-a até o tamanho de itens do slice, por exemplo:

```go
func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{
        "https://random-status-code.herokuapp.com/", 
        "https://www.alura.com.br", 
        "https://www.caelum.com.br",
    }

    for i := 0; i < len(sites); i++ {
        fmt.Println(sites[i])
    }

    site := "https://www.alura.com.br"
    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}
```

Mas há uma forma mais enxuta de fazer isso em Go, utilizando o `range`. Ela é como se fosse um operador de iteração do Go, nos dando acesso a cada item do array, ou do slice, e ele nos retorna dois valores: a posição do item iterado e o próprio item daquela posição:

```go
func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{
        "https://random-status-code.herokuapp.com/", 
        "https://www.alura.com.br", 
        "https://www.caelum.com.br",
    }

    for i, site := range sites {
        fmt.Println("Estou passando na posição", i,
            "do meu slice e essa posição tem o site", site)
    }

    site := "https://www.alura.com.br"
    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}
```

Agora que conseguimos passar por cada item do slice, basta fazer com cada um deles o que já fazemos fora do `for`, ou seja, fazer uma requisição para o site e verificar o seu status code.

Faremos isso no próximo vídeo.





----


# A Instrução `for` em Go

## 📚 **Resumo da Aula**

Esta aula aborda como percorrer slices em Go utilizando a estrutura de repetição `for`, demonstrando duas abordagens diferentes para iterar sobre coleções.

---

## 🔄 **Evolução do Código**

### Alterações Implementadas
- **Remoção**: Função `exibeNomes()` (usada para aprender slices)
- **Implementação**: Slice de sites na função `iniciarMonitoramento()`

### Novo Slice de Sites
```go
sites := []string{
    "https://random-status-code.herokuapp.com/", 
    "https://www.alura.com.br", 
    "https://www.caelum.com.br"
}
```

---

## 🔁 **Formas de Percorrer Slices**

### 1. For Tradicional
```go
for i := 0; i < len(sites); i++ {
    fmt.Println(sites[i])
}
```
- ✅ Usa índice incremental
- ✅ Controle manual do limite com `len(sites)`
- ✅ Acesso ao elemento via `sites[i]`

### 2. For com Range (⭐ Recomendado)
```go
for i, site := range sites {
    fmt.Println("Estou passando na posição", i,
        "do meu slice e essa posição tem o site", site)
}
```

---

## 🔍 **Range - Operador de Iteração**

### Características Principais
- **Retorna dois valores**: índice (`i`) e valor do elemento (`site`)
- **Sintaxe limpa**: Mais legível que o for tradicional  
- **Idiomático em Go**: Forma preferida para iterar coleções

### Vantagens do Range
1. **Menos propenso a erros**: Não há risco de `index out of bounds`
2. **Código mais limpo**: Elimina necessidade de `len()` e incremento manual
3. **Flexibilidade**: Pode ignorar valores não utilizados

---

## 💡 **Dicas Práticas**

### Ignorando Valores Desnecessários
```go
// Se não precisar do índice, use underscore:
for _, site := range sites {
    // trabalhar apenas com site
}

// Se não precisar do valor, use apenas o índice:
for i := range sites {
    // trabalhar apenas com índice i
}
```

---

## 🎯 **Próximos Passos**

**Objetivo**: Integrar verificação de status HTTP dentro do loop range para monitorar todos os sites automaticamente.

**Implementação**: Mover a lógica de requisição HTTP para dentro da estrutura de repetição, aplicando a verificação para cada site do slice.

---

*📝 **Nota**: Esta aula prepara a base para automatizar o monitoramento de múltiplos sites simultaneamente.*