
# 01 Tipos das variáveis

## Transcrição

Tá na hora de começarmos a construir o nosso monitorador de sites. Começaremos exibindo a mensagem inicial, no estilo "Olá, sr. Douglas", mais a versão do sistema. Começando pelo nome, ele irá variar, e não queremos que ele fique preso na frase, logo vamos extrai-lo para uma variável
Declarando uma variável em Go

Para declarar uma variável em Go, utilizamos a palavra var seguida do nome da variável mais o seu tipo. Como é um texto, o seu tipo será string, e vamos logo inicializá-la:

~~~~go
package main

import "fmt"

func main() {
    var nome string = "Douglas"
    fmt.Println("Olá sr.")
}
~~~~

Agora, para concatenar a frase com a variável, passamos a mesma para a função Println, separando os elementos por vírgula:

~~~~go
package main

import "fmt"

func main() {
    var nome string = "Douglas"
    fmt.Println("Olá sr.", nome)
}
~~~~

Falta agora exibirmos a versão do sistema, que será no modelo 1.0, 1.1, etc. Logo, a versão é um número flutuante, que no Go tem duas versões, o float32 para 32 bits, e float64 para 64 bits. Como o nosso número será pequeno, utilizaremos a versão de 32 bits:

~~~~go
package main

import "fmt"

func main() {
    var nome string = "Douglas"
    var versao float32 = 1.1
    fmt.Println("Olá sr.", nome)
    fmt.Println("Este programa está na versão", versao)
}
~~~~

Variável sem valor atribuído

Do mesmo jeito, podemos exibir a idade, que é um número inteiro, logo um int:

~~~~go
package main

import "fmt"

func main() {
    var nome string = "Douglas"
    var idade int
    var versao float32 = 1.1
    fmt.Println("Olá sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)
}
~~~~

A variável idade está vazia propositalmente, e ao executar o programa, temos a seguinte saída:

~~~~bash
alura@alura-01:~/go/src/hello$ go run hello.go
Olá sr. Douglas sua idade é 0
Este programa está na versão 1.1
~~~~

No Go, quando não atribuímos um valor a uma variável, ela assume um valor zerado, ou seja, se for um inteiro, seu valor será 0, se for um número flutuante, seu valor será 0.0, e se for uma string, seu valor será uma string vazia.
Não utilização de uma variável

Outra característica do Go é não podermos declarar uma variável e não utilizá-la. Por exemplo, se removermos a impressão da idade:

~~~~go
package main

import "fmt"

func main() {
    var nome string = "Douglas"
    var idade int
    var versao float32 = 1.1
    fmt.Println("Olá sr.", nome)
    fmt.Println("Este programa está na versão", versao)
}
~~~~

Temos a seguinte saída:

~~~~bash
alura@alura-01:~/go/src/hello$ go run hello.go
# command-line-arguments
./hello.go:7: idade declared and not used
~~~~

Essa é mais uma convenção do Go, temos sempre que utilizar as variáveis que declaramos. Até por que, se não estamos utilizando a variável, não tem motivo dela estar ali.