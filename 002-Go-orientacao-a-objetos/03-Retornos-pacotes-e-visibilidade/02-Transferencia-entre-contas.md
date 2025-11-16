# 02 - Transferência entre contas

## Transcrição

Com base no que desenvolvemos até aqui podemos sacar ou depositar numa conta corrente. Mas o banco digital para o qual trabalhamos também permitirá transferir dinheiro entre contas.

Criaremos um cenário na função main() com contaDaSilvia e contaDoGustavo. Elas serão do tipo ContaCorrente, os nomes dos titulares estarão entre aspas duplas e o saldo de "Silvia" será "300". O de "Gustavo", "100".

Vamos visualizar ambas as contas no terminal também.

```go
func main() {
    contaDaSilvia := ContaCorrente{titular:"Silvia", saldo: 300}
    contaDoGustavo := ContaCorrente{titular:"Gustavo", saldo: 100}
    
    fmt.Println(contaDaSilvia)
    fmt.Println(contaDoGustavo)
}
```

Vamos salvar e executar esse código com go run main.go. Teremos no console {Silvia 0 0 300} e {Gustavo 0 0 100}, ou seja, as informações sobre o titular e o saldo das contas.

Queremos possibilitar a transferência entre as duas contas Para isso, criaremos uma função chamada Transferir() e ela terá um parâmetro. Com base no que já fizemos, quando queríamos sacar, criamos a variável valorDoSaque. Para depositar, valorDoDeposito. Agora vamos gerar valorDaTransferencia, que será do tipo float64.

Para que uma transferência aconteça precisamos informar que estamos tirando o valor de uma conta e enviando para outra. Criaremos um tipo contaDestino. Na linguagem Go, nós especificamos os tipos dos nossos parâmetros. Pensando na struct que fizemos anteriormente, colocamos a palavra type. Assim, a própria ContaCorrente é um tipo.

Então, podemos dizer que a contaDestino é do tipo ContaCorrente. Além disso, criaremos um parâmetro booleano que servirá como retorno para essa função, dizendo se é verdadeiro ou falso.

Dentro do corpo da função, teremos que verificar se a conta a partir da qual faremos a transferência tem saldo. Se Silvia quiser transferir para Gustavo, por exemplo, precisaremos verificar se a conta dela tem saldo para isso. Colocaremos antes da função um ponteiro para a ContaCorrente que chamará essa função. Faremos um if colocando a condição de que se o valorDaTransferencia for menor do que o saldo da conta chamando, será possível fazer a transferência

Primeiro será necessário remover dinheiro da conta que transferirá. Para isso, vamos tirar do c.saldo o valorDaTransferencia. Depois, vamos ter que enviar o dinheiro para a outra conta, o que será um processo semelhante ao de depositar. Assim, usaremos Depositar como um método para contaDestino receber o valor transferido.

Isso ainda não será suficiente, pois falta informar o retorno booleano. Portanto, se o procedimento acontecer, retornaremos true, ou seja, o retorno da transferência é verdadeiro.

Caso não aconteça, faremos um else para retornar um false

```go
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino ContaCorrente) bool {
    if valorDaTransferencia < c.saldo {
        c.saldo -= valorDaTransferencia 
        contaDestino.Depositar(valorDaTransferencia)
        return true
    } else {
        return false
    }
}
```

Vamos salvar e aparentemente não há problemas, então vamos testar. Sabemos que o saldo de contaDaSilvia é "300" e de contaDoGustavo é "100". Para armazenar o resultado, criaremos uma variável chamada status e faremos a transferência a partir de contaDaSilvia. Colocaremos o valor "200" nos parâmetros de Transferir() , e no segundo parâmetro, o destino, contaDoGustavo.

Se apenas salvarmos a aplicação agora, será avisado que a variável status foi declarada mas não foi utilizada. Então vamos exibi-lo no print e abaixo o conteúdo das duas contas.

```go
status := contaDaSilvia.Transferir(200, contaDoGustavo)

fmt.Println(status)
fmt.Println(contaDaSilvia)
fmt.Println(contaDoGustavo)
```

Vamos executa o código e veremos no terminal que a transferência aconteceu. A conta da Silvia, anteriormente com saldo de "300", após tranferir "200" para Gustavo está com "100".

A conta do Gustavo tinha "100" de saldo. Silvia fez a transferência de "200" para a conta dele, mas ainda aparecerá no console que o saldo dele é "100". Ou seja, algo está errado. O que será que aconteceu?

O valor da conta da Silvia foi retirado, mas o conteúdo não entrou na conta de destino. O erro aconteceu porque fizemos referência para contaDaSilvia, apontamos para o conteúdo da conta que chama a função para alterar o conteúdo dela, mas não fizemos o mesmo para contaDestino.

Como queremos alterar o valor dessa conta também, colocamos um asterisco na frente dela nos parâmetros da função.

```go
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
    if valorDaTransferencia < c.saldo {
        c.saldo -= valorDaTransferencia 
        contaDestino.Depositar(valorDaTransferencia)
        return true
    } else {
        return false
    }
}
```

Porém, se salvarmos essa alteração, será apresentado um erro, e o alerta nos dirá que estamos apontando para um local desconhecido. Precisamos identificar contaDoGustavo. Conseguimos fazer isso colocando um & quando fizermos a transferência. Assim, dizemos que queremos transferir de fato para esse endereço além de termos apontado para a conta.

```go
status := contaDaSilvia.Transferir(200, &contaDoGustavo)
```

Faremos um teste, pressionaremos "Ctrl + L" e executaremos. Agora sim, a conta de Silvia que tinha "300" como saldo passará a ter "100" e a de Gustavo, "300", pois recebeu "200".

Vamos fazer o contrário, pegar a contaDoGustavo e transferir "200" para Silvia, levando em conta que ele não tem saldo suficiente para isso, pois o saldo dele é "100".

```go
status := contaDoGustavo.Transferir(200, &contaDaSilvia)
```

Veremos no console a mensagem de falso para a transferência, ela não aconteceu, e os saldos das contas permanecerão os mesmos.

Vamos ver o que acontecerá se Gustavo tentar transferir "-200" para a conta de Silvia.

```go
status := contaDoGustavo.Transferir(-200, &contaDaSilvia)
```

Vamos pressionar "Ctrl + L", executar e veremos a mensagem "true", de que foi verdadeiro. A conta de Silvia continuará com "300", mas a de Gustavo também terá "300". Esse é um comportamento que não esperávamos.

Isso significa que alem de verificar se o valor da transferência é menor do que o valor do saldo, verificaremos também se o valor da transferência é maior do que 0. As duas condições precisam ocorrer simultaneamente para que a transferência ocorra e isso seja verdade.

```go
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

Vamos salvar, limpar o terminal e executar e dessa vez a transferência não aconteceu, pois o resultado será falso e o saldo das contas permanecerá o mesmo do início, "300" para Silvia e "100" para Gustavo.