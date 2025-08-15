# 09  Mãos na Massa: Organização e Fluxo  

Começando deste ponto? Você pode fazer o DOWNLOAD completo do projeto do capítulo anterior e continuar seus estudos a partir deste capítulo.

Neste capítulo vamos determinar para qual fluxo do nosso programa o usuário deve seguir a partir do comando que ele escolheu, e além disso vamos começar a organizar nosso código em pequenas funções.

## 1) Crie um switch com as opções do menu e um caso default

```go
// hello.go
package main

import "fmt"

func main() {
    // ... restante da função main
    var comando int
    fmt.Scan(&comando)

    fmt.Println("O valor da variável comando é:", comando)

    switch comando {
    case 1:
        fmt.Println("Monitorando...")
    case 2:
        fmt.Println("Exibindo Logs...")
    case 0:
        fmt.Println("Saindo do programa...")
    default:
        fmt.Println("Não conheço este comando")
    }
}
```

_Aqui colocamos uma opção no switch para cada item do menu que tínhamos anteriormente._

## 2) Extraia a introdução para a função exibeIntroducao

```go
// hello.go
func main() {
    exibeIntroducao()

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    // ...  restante da função main
}

func exibeIntroducao() {
    nome := "Douglas"
    versao := 1.1
    fmt.Println("Olá, sr(a).", nome)
    fmt.Println("Este programa está na versão", versao)
}
```

_Não esqueça de chamar a função exibeIntroducao() na função main._

## 3) Extraia o menu para a função exibeMenu

```go
func main() {
    exibeIntroducao()
    exibeMenu()

    var comando int
    fmt.Scan(&comando)
    fmt.Println("O valor da variável comando é:", comando)
    // ... restante da função
}

func exibeIntroducao() {
    // ... restante da função
}

func exibeMenu() {
    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")
}
```

_Não esqueça de chamar a função exibeMenu() na função main._

## 4) Crie a função leComando que retorna um int

```go
func main() {
    exibeIntroducao()
    exibeMenu()
    comando := leComando()

    switch comando {
    // ... restante da função
    }
}

func exibeIntroducao() {
    // ... restante da função
}

func exibeMenu() {
    // ... restante da função
}

func leComando() int {
    var comandoLido int
    fmt.Scan(&comandoLido)
    fmt.Println("O valor da variável comando é:", comandoLido)

    return comandoLido
}
```

_Não esqueça de capturar o comando lido na função main, para que ele possa ser utilizado pelo switch._

## 5) Adicione os pontos de saída (os.Exit)

```go
package main

import "fmt"
import "os"

func main() {
    exibeIntroducao()
    exibeMenu()
    comando := leComando()

    switch comando {
    case 1:
        fmt.Println("Monitorando...")
    case 2:
        fmt.Println("Exibindo Logs...")
    case 0:
        fmt.Println("Saindo do programa...")
        os.Exit(0)
    default:
        fmt.Println("Não conheço este comando")
        os.Exit(-1)
    }
}

// ... restante do arquivo
```


