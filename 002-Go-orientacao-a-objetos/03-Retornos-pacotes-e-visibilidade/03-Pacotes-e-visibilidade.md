# 03 Pacotes e visibilidade

## Transcrição

Até esse momento do curso, nosso código conta com 50 linhas. A medida que novas funcionalidades são implementadas, mais linhas de código surgem e pode se tornar mais difícil a manutenção dele.

É importante lembrar que não desenvolvemos um código só para nós entendermos, outros desenvolvedores também precisam entendê-lo para poder trabalhar nele. Por isso devemos mantê-lo de uma maneira simplificada.

Mas como podemos torná-lo mais organizado sem alterar todo o conteúdo e os comportamentos que já criamos?

Para isso, podemos utilizar uma convenção do Go que se trata da distribuição do nosso código em pacotes. Vamos lembrar que o primeiro comando utilizado no Go é informar qual é o pacote que utilizamos, `package main`.

Então, podemos criar um pacote de contas e colocar nele os trechos referentes às Contas Correntes. Caso surjam ainda outras contas, como a Poupança, elas também serão inseridas no pacote.

Clicaremos no primeiro link da barra na lateral esquerda do Visual Studio Code, o botão "Explorer" ou "Ctrl + Shift + E". Todo nosso código estará nesse local. Criaremos uma pasta chamada "contas" nessa área, e dentro dela criaremos um "New file" (novo arquivo) chamado "contaCorrente.go", porque manteremos aqui somente os códigos de conta corrente.

Nesse arquivo, usaremos o comando `package contas`. Todo nosso conteúdo de conta corrente deverá ficar dentro desse pacote. Por isso vamos usar "Ctrl + X" para trazer os trechos do código necessários para dentro do projeto.

```go
package contas

type ContaCorrente struct {
    titular       string
    numeroAgencia int
    numeroConta   int
    saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
    podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
    if podeSacar {
        c.saldo -= valorDoSaque
        return "Saque realizado com sucesso"
    } else {
        return "Saldo insuficiente"
    }
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
    if valorDoDeposito > 0 {
        c.saldo += valorDoDeposito
        return "Deposito realizado com sucesso", c.saldo
    } else {
        return "Valor do deposito menor que zero", c.saldo
    }
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
    if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
        c.saldo -= valorDaTransferencia
        contaDestino.Depositar(valorDaTransferencia)
        return true
    } else {
        return false
    }
}
```

Vamos salvar e não teremos nenhum problema. Vamos voltar ao main e tentaremos salvar também, porém ele apresentará uma mensagem de erro que diz que ContaCorrente está indefinida. Colocamos ContaCorrente no nosso pacote de contas, então precisamos avisar para o main em que local ele pode encontrar a informação.

Na nossa primeira linha de main.go fazemos o import de uma classe para conseguirmos visualizar a saída do nosso desenvolvimento no terminal, o pacote "fmt". Faremos outro import colocando parênteses e todos os imports dentro deles.

Podemos descobrir o caminho para nosso pacote de contas no terminal. Vamos limpar nosso terminal e depois digitaremos `pwd`. Aparecerá no console todo o caminho das pastas do sistema até onde estarão os pacotes. No nosso caso, poderemos copiar a partir de `github.com/alura/banco`. Entretanto, se digitarmos só até "bancos", ele não entrará na pasta contas, por isso estenderemos o caminho até ela.

```go
package main

import (
    "fmt"
    "github.com/alura/banco/contas"
)
```

Porém se se salvarmos, o caminho que digitamos sumirá. Isso acontece porque trouxemos o pacote de contas, mas ainda não dissemos qual informação contida nele foi utilizada no arquivo em que estamos trabalhando. Para identificar que a ContaCorrente em main.go é a mesma de contas.go, precisamos dizer de onde ContaCorrente veio, escrevendo `contas.` antes de utilizá-la.

```go
package main

import (
    "fmt"

    "github.com/alura/banco/contas"
)

func main() {
    contaDaSilvia := contas.ContaCorrente{titular: "Silvia", saldo: 300}
    contaDoGustavo := contas.ContaCorrente{titular: "Gustavo", saldo: 100}

    status := contaDoGustavo.Transferir(-200, &contaDaSilvia)

    fmt.Println(status)
    fmt.Println(contaDaSilvia)
    fmt.Println(contaDoGustavo)
}
```

Agora vamos salvar e nosso código estará correto nesse ponto. Mas ainda haverá um outro erro. Ele está nos alertando de que não sabe o que é o campo titular nem o saldo por questões de visibilidade.

Visibilidade é um atributo referente a uma função ou variável referente a ser visível por outras partes da aplicação, como os pacotes, de forma muito semelhante ao private ou public do Java.

Não temos nenhum erro no nosso código de contas. Nele, deixamos todas as primeiras letras dos nossos campos em minúsculas. Com isso, tornamos eles visíveis apenas por esse arquivo.

> **Importante!**: Na linguagem Go, se escrevermos nomes de variáveis deixando a primeira letra minúscula, elas ficarão visíveis apenas no arquivo em que foram declaradas.

Para alterar a visibilidade dos campos, alteraremos todas as primeiras letras dos nomes para maiúsculas, tanto nas declarações quanto nos trechos em que titular e saldo são utilizados.

Para encontrar as palavras em todos os pontos do código, é possível pressionar "Ctrl + F" e procurá-las para então fazer a substituição. Após encontrar os termos, clicaremos na seta à direita da janela de busca e o campo "Replace" será aberto. Nele escreveremos a forma para a qual queremos alterar, por exemplo, buscando por "saldo", escreveremos "Saldo". Faremos essa modificação para todas as vezes que as palavras aparecem clicando no ícone "Replace all".

```go
package contas

type ContaCorrente struct {
    Titular       string
    NumeroAgencia int
    NumeroConta   int
    Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
    podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
    if podeSacar {
        c.Saldo -= valorDoSaque
        return "Saque realizado com sucesso"
    } else {
        return "Saldo insuficiente"
    }
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
    if valorDoDeposito > 0 {
        c.Saldo += valorDoDeposito
        return "Deposito realizado com sucesso", c.Saldo
    } else {
        return "Valor do deposito menor que zero", c.Saldo
    }
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
    if valorDaTransferencia < c.Saldo && valorDaTransferencia > 0 {
        c.Saldo -= valorDaTransferencia
        contaDestino.Depositar(valorDaTransferencia)
        return true
    } else {
        return false
    }
}
```

Então colocamos todos os códigos referentes a Conta Corrente, a struct e todos os nossos métodos num pacote de contas e caso outras contas apareçam, elas podem ser colocadas na mesma pasta, seja por quem precisar modificar a aplicação. Se for necessário usar uma nova conta adicionada, a prática será fazer um import dela também.

Também podemos dar um outro nome para nosso import. Por exemplo, escrevendo uma letra "c" antes de passar o caminho, c será entendido como um apelido.

```go
import (
    "fmt"

    c "github.com/alura/banco/contas"
)
```

Assim, quando formos escrever o pacote de onde vem uma informação, podemos utilizar só o c antes do nome do tipo, na nossa situação.

```go
contaDaSilvia := c.ContaCorrente{titular: "Silvia", saldo: 300}
```

Se salvarmos, teremos o mesmo resultado. Mas manteremos a escrita `contas.ContaCorrente` para não confundir. Agora já sabemos que podemos dar apelidos aos nossos imports.



## RESUMO

Nesta aula, aprendemos sobre a organização do código Go em pacotes para facilitar a manutenção e o entendimento por outros desenvolvedores.

Criação de Pacotes: Demonstramos como criar um pacote chamado "contas" para agrupar códigos relacionados a contas correntes.

Visibilidade: Abordamos o conceito de visibilidade em Go, onde campos e funções com a primeira letra minúscula são visíveis apenas dentro do mesmo pacote, enquanto os com a primeira letra maiúscula são visíveis externamente.

> **Importante!**: Na linguagem Go, se escrevermos nomes de variáveis deixando a primeira letra minúscula, elas ficarão visíveis apenas no arquivo em que foram declaradas.

Para alterar a visibilidade dos campos, alteraremos todas as primeiras letras dos nomes para maiúsculas, tanto nas declarações quanto nos trechos em que titular e saldo são utilizados.
Foram passados para Maiusculo a primeira Letra de cada item da Struct:
```go
package contas

type ContaCorrente struct {
    Titular       string
    NumeroAgencia int
    NumeroConta   int
    Saldo         float64
}
// RESTANTE DO CÓDIGO OMITIDO
```

No código do main.go também é necessário passar para Maiusculo onde for usar:

~~~~go
package main

import (
// RESTANTE DO CÓDIGO OMITIDO
)

func main() {
	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}
// RESTANTE DO CÓDIGO OMITIDO
}
~~~~

Importação de Pacotes: Mostramos como importar pacotes em outros arquivos Go e como usar um alias (apelido) para referenciar o pacote importado.

