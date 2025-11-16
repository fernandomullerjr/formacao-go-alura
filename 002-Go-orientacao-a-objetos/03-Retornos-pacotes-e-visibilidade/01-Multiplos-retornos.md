# 01 - Múltiplos retornos

## Transcrição

Criamos um método para sacar, que verifica se o valor do saque é maior do que 0 e se ele é maior do que o valor do saldo que um cliente possui na conta.

Entretanto, podemos sacar no banco, mas ainda não podemos depositar. Criaremos uma função que permita depositar dinheiro na conta. A função Depositar() receberá como parâmetro o valor do depósito, que vamos pegar e atribuir ao saldo.

Portanto, nos parênteses de Depositar() estará valorDoDeposito, do tipo float 64. Se uma outra conta chamar nossa função, devemos apontar para ela, então colocaremos antes da função c apontando para a conta corrente que estiver chamando.

Nossa função pode ser bem simples. Queremos que ela pegue o saldo da conta que estiver chamando c.saldo e atribua o valor do depósito. Usamos o -= para remover o valor do saque do saldo. Agora usaremos += para conceder o valor do depósito.

Já fizemos o teste na função main() com o saque. Agora faremos para Depositar().

```go
func (c *ContaCorrente) Depositar(valorDoDeposito float64) {
    c.saldo += valorDoDeposito
}

func main() {
    contaDaSilvia := ContaCorrente{}
    contaDaSilvia.titular = "Silvia"
    contaDaSilvia.saldo = 500
    
    fmt.Println(contaDaSilvia.saldo) 
}
```

Vamos salvar e executar o código e pressionar "Command + J" para visualizar no terminal. Rodaremos a aplicação com go run main.go e veremos o saldo de "500".

Agora vamos atribuir o valor do depósito. Para isso, no fim do código escreveremos contaDaSilvia.Depositar(). Quando abrimos os parênteses o valorDoDeposito já será solicitado. Vamos depositar mais "500" à conta. Na sequência, faremos um print para exibir novamente o valor do saldo.

```go
fmt.Println(contaDaSilvia.saldo) 
contaDaSilvia.Depositar(500)

fmt.Println(contaDaSilvia.saldo) 
```

Mais uma vez vamos limpar o terminal e executar nosso código. Veremos que tínhamos o valor de "500" e após o depósito chegamos em "1000".

Será que teremos problemas com números negativos? Temos "500" de saldo na conta de Silvia. Depositaremos "-1000" para entender o que acontecerá:

```go
contaDaSilvia.Depositar(-1000)

fmt.Println(contaDaSilvia.saldo) 
```

Agora, ficamos com "-500" de saldo. Tínhamos "500" e pedimos para depositar "-1000". Sendo assim, foram subtraídos os "500" que tínhamos e mais "500", ficando um valor negativo. Não é o que queremos que aconteça.

Precisaremos tratar o valor do depósito fazendo a verificação se esse é um valor maior do que 0. Criaremos um if na função para valorDoDeposito > 0. Se o valor do depósito for maior do que 0, aí sim atribuiremos o valor do depósito ao saldo.

```go
func (c *ContaCorrente) Depositar(valorDoDeposito float64) {
    if valorDoDeposito > 0 {
        c.saldo += valorDoDeposito
    }
}
```

Vamos salvar e fazer o teste novamente para o depósito de "-1000". Rodaremos o código e teremos "500" e "500" de saldo impresso, ou seja, ele não se alterou porque nenhum depósito foi efetuado. Mesmo colocando um valor negativo bastante alto, o depósito não ocorrerá.

Mas se depositarmos um valor positivo, ele será somado. Por exemplo, se fizermos contaDaSilvia.Depositar(1000), o saldo, por fim será "1500".

Mas nossa função de depósito não nos retorna nada, ela simplesmente verifica e já envia os valores, diferentemente da nossa função de saque em que uma mensagem nos informa se ele foi efetuado com sucesso ou não.

Vamos alterar o código adicionando uma string como retorno quando o depósito é realizado. Caso o valor depositado não seja maior do que 0 e a ação não seja efetuada, a string avisará "Valor depósito menor que zero".

Para mostrar essa mensagem no console, faremos fmt.Println(contaDaSilvia.Depositar(2000)) ao fim.

```go
func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {
    if valorDoDeposito > 0 {
        c.saldo += valorDoDeposito
        return "Deposito realizado com sucesso"
    } else { 
        return "Valor do depósito menor que zero"
    }
}

func main() {
    contaDaSilvia := ContaCorrente{}
    contaDaSilvia.titular = "Silvia"
    contaDaSilvia.saldo = 500
    
    fmt.Println(contaDaSilvia.saldo) 
    fmt.Println(contaDaSilvia.Depositar(2000)) 
}
```

Vamos salvar e executar a função e receberemos a mensagem de que o depósito foi realizado no terminal. Mas além de saber que o depósito foi feito, queremos saber o novo saldo da conta. Será que conseguimos criar uma função que nos devolva mais de um retorno?

Usando Go conseguimos fazer isso se, primeiramente, adicionarmos (string, float64) após (c *ContaCorrente) Depositar(valorDoDeposito float64). Assim, diremos que nossa função apontará sempre para a ContaCorrente que a chamar, tem um parâmetro do que é o valorDoDeposito e devolverá uma mensagem de string e um float, o saldo. Após as duas mensagens colocaremos, portanto, o c.saldo como retorno também.

```go
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
    if valorDoDeposito > 0 {
        c.saldo += valorDoDeposito
        return "Deposito realizado com sucesso", c.saldo
    } else { 
        return "Valor do depósito menor que zero", c.saldo
    }
}
```

Vamos rodar novamente, ainda para o depósito de "2000". Haverá no terminal o saldo inicial de "500", a mensagem de que o depósito foi realizado e na sequência o valor "2500", saldo final após a atribuição do valor.

Assim, aprendemos mais um ponto importante na linguagem Go. Observe que podemos ter um retorno, como quando criamos o saque, nenhum, como quando fizemos pela primeira vez o comportamento de depositar e nada era devolvido, e mais de um retorno como no último caso em que trabalhamos.

Também conseguimos trabalhar com os retornos. Por exemplo, criaremos uma variável chamada status e outra chamada valor. Atribuiremos nessa variável o resultado desse depósito.

O que fizemos foi: como chamamos essa função, ela nos devolverá dois itens. Então, o primeiro retorno dela, uma mensagem string, será atribuída ao status. O segundo, o novo saldo, caso seja realizado com sucesso vai para a segunda variável, valor.

Sendo assim, printaremos as duas variáveis.

```go
fmt.Println(contaDaSilvia.saldo) 
status, valor := contaDaSilvia.Depositar(2000)
fmt.Println(status, valor) 
```

Vamos salvar e executar o programa. Teremos ainda o mesmo resultado. Podemos ter uma função com mais de um retorno e trabalhar esses retornos de forma individual também.




---



## Dia 15/11/2025

- Começando a aula com base no código da aula abaixo:
formacao-go-alura/002-Go-orientacao-a-objetos/02-Referencia-ponteiro-metodos/03-metodo-sacar.go

usando o código que tem a parte de saque:

~~~~go
package main

import (
	"fmt"
)

// ContaCorrente representa uma conta bancária simples.
// Os campos estão em minúsculas, portanto são privados ao pacote atual.
type ContaCorrente struct {
	titular       string  // nome do titular da conta
	numeroAgencia int     // número da agência bancária
	numeroConta   int     // número da conta
	saldo         float64 // saldo disponível
}

// Método "Sacar" — responsável por diminuir o saldo, caso haja valor suficiente.
//
// Receiver: *ContaCorrente (ponteiro)
// → Permite alterar o valor original de 'saldo' dentro da struct.
//
// Parâmetro: valorDoSaque (float64)
// → Quantia que o cliente deseja sacar.
//
// Retorna: string — mensagem de sucesso ou erro.
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	// Regra: só permite saque se o valor for positivo e <= saldo atual
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		// Atualiza o saldo (modifica o valor original, pois c é um ponteiro)
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {
	// Cria uma instância (objeto) de ContaCorrente com zero values.
	contaDaSilvia := ContaCorrente{}

	// Define manualmente alguns campos.
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	// Mostra o saldo inicial.
	fmt.Println(contaDaSilvia.saldo) // 500

	// Tenta sacar 400 reais — o método altera o saldo original.
	fmt.Println(contaDaSilvia.Sacar(400)) // "Saque realizado com sucesso"

	// Exibe o novo saldo após o saque.
	fmt.Println(contaDaSilvia.saldo) // 100
}

~~~~



- Ajustando o código.
- Efetuando teste:

> cd /home/fernando/cursos/golang/formacao-go-alura/002-Go-orientacao-a-objetos/03-Retornos-pacotes-e-visibilidade
> ls
01-Multiplos-retornos.md  01-multiplos-retornos.go  go_oo-aula3
> go run 01-multiplos-retornos.go
500
Saque realizado com sucesso
100
100
Deposito realizado com sucesso 2100
> date
Sat Nov 15 22:31:53 -03 2025