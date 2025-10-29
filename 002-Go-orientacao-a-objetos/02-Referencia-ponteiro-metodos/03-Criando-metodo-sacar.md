# 03 Criando o método sacar

## Transcrição

Sacar ou depositar um dinheiro em uma conta corrente é um comportamento comum.

Com base no código que desenvolvemos até agora desenvolveremos um comportamento de saque. Removeremos as nossas comparação entre tipos e criaremos uma nova conta corrente, `contaDaSilvia`.

Ela será do tipo `ContaCorrente`, terá como titular a string "Silvia" e saldo de 500. Vamos exibir a conta para visualizá-la no nosso terminal.

````go
func main() {
    contaDaSilvia := ContaCorrente{}
    contaDaSilvia.titular = "Silvia"
    contaDaSilvia.saldo = 500

    fmt.Println(contaDaSilvia)
}
````

Executamos e no terminal haverá `{Silvia 0 0 500}`, ou seja, o nome do titular, dois zeros referentes aos valores de Número da Agência e Conta que não adicionamos, e "500" que é o valor do nosso saldo. Podemos exibir apenas o saldo:

````go
fmt.Println(contaDaSilvia.saldo)
````

Executaremos mais uma vez e veremos apenas o valor do saldo, "500".

Para sacar um valor, criaremos uma nova variável, `valorDoSaque`, e usaremos dois pontos igual (`:=`) para atribuir um valor de "200". Se a conta tem "500" e removeremos "200", ficaremos com "300" de saldo. Então o valor do saldo, `contaDaSilvia.saldo` será igual a `contaDaSilvia.saldo` menos `valorDoSaque`. Depois, exibiremos mais uma vez o valor que Silvia possui na conta.

````go
valorDoSaque := 200
contaDaSilvia.saldo = contaDaSilvia.saldo - valorDoSaque

fmt.Println(contaDaSilvia.saldo)
````

Quando salvamos, porém, o compilador mostra um alerta de erro, e se passarmos o mouse sobre a linha sublinhada em vermelho entenderemos que o problema é tentamos fazer uma operação entre um valor do tipo float e outro do tipo inteiro. `valorDoSaque` foi criado como um inteiro, mas o saldo é um float. Nesse caso, bastará colocar um ponto (`.`) após o valor atribuído ao `valorDoSaque` e a linguagem entenderá que ele é um float.

````go
valorDoSaque := 200.
````

Agora, quando executarmos o código, teremos primeiro o "500", valor de saldo que printamos, e depois o cálculo será efetuado e veremos como resultado o valor "300" no terminal. Mas vamos supor que queremos remover "800" da conta. A cliente não tem esse saldo, então ultrapassaremos o valor.

````go
valorDoSaque := 800.
````

Vamos rodar e o que aparecerá no terminal agora será "-300". Isso não deveria acontecer, pois temos apenas "500" na conta. Para evitar o problema podemos gerar uma verificação para esse saque.

Pensando nisso, criaremos uma função que verificará se o valor que tentaremos sacar será menor do que o valor presente na conta. Vamos criá-la antes da `func main()`, logo após nossa struct, chamaremos a função de `Sacar()` e ela receberá um valor. Nosso valor do saque deverá ser executado nos parênteses da função. Também especificaremos que ele será do tipo `float64`.

Depois dos parâmetros, retornaremos uma mensagem do tipo string nesse campo, informando se o saque foi realizado com sucesso ou não. Então criaremos uma variável para verificar se é possível sacar com `podeSacar :=`, pegaremos `valorDoSaque` e vamos comparar se ele é menor ou igual ao saldo de `contaDaSilvia`.

Podemos fazer `podeSacar := valorDoSaque <= contaDaSilvia.saldo`, porém nesse campo precisamos dizer que quem tentará fazer o saque será a pessoa responsável pela conta. Nesse caso só temos `contaDaSilvia`, mas se tivermos que adicionar ainda outros clientes, teremos criado uma função que só funciona para `contaDaSilvia`. Queremos que sempre quando alguém tentar sacar, apontemos para a conta dessa pessoa, de forma semelhante ao que acontece com o `this` do Java ou o `self` do Python.

Para referenciarmos esse ponteiro no momento da criação do tipo, podemos colocar entre parênteses logo após `func` e antes de `Sacar()` a inscrição `(c *ContaCorrente)`, o que significa que quando a função for chamada, o código apontará para a conta que chama. Nesse caso, quando chamarmos a função, não precisaremos especificar que nos tratamos da conta de um cliente ou outro.

Nesse caso, se a conta que estiver chamando a função tiver saldo, será possível sacar. Criaremos uma condicional `if` para fazer a verificação. Poderemos sacar se for verdadeiro que `valorDoSaque` é menor do que saldo. Se podemos sacar colocaremos `c.saldo` no corpo do `if`. Podíamos escrever `conta.saldo`, se escrevêssemos `(conta *ContaCorrente)`, mas por uma questão da linguagem Go, sempre utilizamos a primeira letra do nome no nosso ponteiro dentro da função.

Porém, escrevemos muito para fazer nossa operação usando `contaDaSilvia.saldo = contaDaSilvia.saldo - valorDoSaque`. O que queremos fazer é pegar o que há na conta e subtrair pelo valor do saque. Há outra forma de fazer o mesmo, e a utilizaremos na função, com `c.saldo -= valorDoSaque`.

O `valorDoSaque` será o valor que passaremos para esse método. Assim, será possível remover o valor para sacar se ele for menor do que o saldo, e retornaremos uma mensagem do tipo string, entre aspas duplas, "Saque realizado com sucesso".

Mas e se isso não for verdade? Então, teremos o `else` para o caso do valor que tentarmos sacar ser maior do que o saldo. Retornaremos "Saldo insuficiente".

````go
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
    podeSacar := valorDoSaque <= c.saldo
    if podeSacar {
        c.saldo -= valorDoSaque
        return "Saque realizado com sucesso"
    } else {
        return "Saldo insuficiente"
    }
}
````

Criamos função para verificar se o valor do saque é menor do que o do saldo. Se sim, poderemos sacar.

Apagaremos `contaDoGustavo` do fim do código, pois nosso interesse nesse momento é visualizar apenas a conta da Silvia. O valor do saldo dela, "500", já estará sendo exibido com um print.

Agora, no lugar de fazer toda a operação em que o saldo é igual ao saldo menos o saque e vamos apenas chamar a função `Sacar()`. Bastará digitarmos `contaDaSilvia` e um ponto (`.`) para que `Sacar()` seja sugerido.

Como essa função devolve uma string, podemos colocá-la dentro da string dos parâmetros do print. Quando digitamos `contaDaSilvia.`, conseguimos visualizar todos os campos que temos pressionando "Ctrl + espaço" para ter acesso ao atalho. Após `contaDaSilvia.Sacar` terá que vir o abrir e fechar de parênteses `()`, pois queremos executar. Colocando os parênteses, já será recomendado definir o valor do saque. Colocaremos "300".

````go
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
    podeSacar := valorDoSaque <= c.saldo
    if podeSacar {
        c.saldo -= valorDoSaque
        return "Saque realizado com sucesso"
    } else {
        return "Saldo insuficiente"
    }
}

func main() {
    contaDaSilvia := ContaCorrente{}
    contaDaSilvia.titular = "Silvia"
    contaDaSilvia.saldo = 500

    fmt.Println(contaDaSilvia.saldo)

    fmt.Println(contaDaSilvia.Sacar(300))
}
````

Abriremos o terminal com "Command + J", o limparemos com "Ctrl + L" e executaremos mais uma vez.

Veremos o valor "500" que havia na conta e em seguida a mensagem de que o saque foi realizado com sucesso. Pediremos para que o código nos mostre quanto restou na conta da Silvia depois do saque colocando outro print embaixo do saque.

````go
fmt.Println(contaDaSilvia.saldo)

fmt.Println(contaDaSilvia.Sacar(300))
fmt.Println(contaDaSilvia.saldo)
````

Agora veremos que tínhamos "500", realizamos com sucesso o saque (de 300) e será impresso que ficamos com "200". Então nossa função terá funcionado para quando temos saldo. Mas e quando não temos?

Tentaremos sacar "600" para testar.

````go
fmt.Println(contaDaSilvia.saldo)

fmt.Println(contaDaSilvia.Sacar(600))
fmt.Println(contaDaSilvia.saldo)
````

Digitaremos "Ctrl + L" e executaremos mais uma vez o programa. Agora veremos a mensagem "Saldo insuficiente" e o valor do saldo permanecerá em "500". Parece que tudo está funcionando, exceto por um detalhe.

Tínhamos "500" de saldo. Ao tentar sacar "600", não foi possível porque "600" é maior do que "500". Só que se o valor do saque for um valor negativo?

Tentaremos sacar "-100".

````go
fmt.Println(contaDaSilvia.saldo)

fmt.Println(contaDaSilvia.Sacar(-100))
fmt.Println(contaDaSilvia.saldo)
````

Quanto teremos de saldo quando executarmos a função? "-100" é menor do que "500", condição para que o saque seja realizado.

Quando rodarmos, veremos que a mensagem é a de que o saque de "-100" foi realizado com sucesso, e o valor restante será "600". Ou seja, a operação efetuada foi a do valor do saldo menos menos (- (-)) o `valorDoSaque`, de forma que o valor ficou positivo e foi atribuído mais "100" ao que havia na conta.

Então, além de verificar se o valor do saque é maior do que o saldo, precisamos simultaneamente verificar se ele é um valor positivo, maior do que 0. Para isso, colocaremos `&&` e assim para que `podeSacar`, que é um booleano (pode ser verdadeiro ou falso), seja verdadeiro, estas duas condições precisam ser verdadeiras.

````go
podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
````

Vamos salvar e executar mais uma vez. Veremos que o saldo era "500", tentando sacar "-100" o saldo será insuficiente e serão mantidos os "500" na conta de Silvia. Não foi possível realizar o saque.

Agora se tentarmos sacar outro valor, como "400", o saque será realizado com sucesso e ainda haverá "100" na conta.
.
````go
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

func main() {
    contaDaSilvia := ContaCorrente{}
    contaDaSilvia.titular = "Silvia"
    contaDaSilvia.saldo = 500

    fmt.Println(contaDaSilvia.saldo)

    fmt.Println(contaDaSilvia.Sacar(400))
    fmt.Println(contaDaSilvia.saldo)
}
````

Este trecho em Go:

```go
c.saldo -= valorDoSaque
```

faz uma **subtração e atribuição**. Ele diminui o valor da variável `c.saldo` pelo valor de `valorDoSaque`. É equivalente a:

```go
c.saldo = c.saldo - valorDoSaque
```

**Explicação:**
- `c` provavelmente é uma instância de uma struct (por exemplo, uma conta bancária).
- `saldo` representa o saldo atual.
- `valorDoSaque` é o valor que está sendo sacado.

**Resumo:**  
Após essa linha, o saldo da conta (`c.saldo`) será reduzido pelo valor do saque (`valorDoSaque`). Isso é comum em métodos que processam saques em sistemas bancários.

**Gotcha:**  
Se não houver validação antes, pode permitir saldo negativo. É importante garantir que o saque só aconteça se houver saldo suficiente.