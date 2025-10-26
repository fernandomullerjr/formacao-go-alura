# 04 Utilizando a struct

## Transcrição

Quando não passamos nenhuma informação para a struct, ela é criada automaticamente com os valores padrão. Mas não é o que queremos, nossa ideia é criar uma conta corrente com todos os campos correspondentes.

Primeiro vamos lembrar que a primeira palavra que usamos para declarar nossa struct é type, assim como os tipos que declaramos para as variáveis serem strings ou inteiros.

Agora criaremos uma nova variável na func main(). Para isso precisaremos alocar um espaço na memória por meio de var e chamaremos a variável de contaDoGuilherme. Ela não será de um tipo inteiro, string ou float, mas do tipo contaCorrente. Por esse motivo utilizamos a palavra type na declaração de contaCorrente, uma estrutura que tem vários campos.

Vamos exibir contaDoGuilherme no terminal na sequência.

````go
func main() {
    var contaDoGuilherme ContaCorrente = ContaCorrente{}

    fmt.Println(contaDoGuilherme)
}
````

Se eu pedir para exibir a conta no terminal, após salvar o código, limpar o terminal e executar mais uma vez, veremos o mesmo resultado de antes, um zero value ou a inicialização automática com valores que o Go coloca.

Mas a forma de escrita não está tão adequada a que costuma-se utilizar na linguagem Go. Quando queremos alocar algo na memória e já fazer uma atribuição, podemos fazer uma atribuição curta. No lugar de usar o sinal de igual, colocaremos dois pontos antes do símbolo, apagaremos o tipo contaCorrente e poderemos apagar também a palavra var. Já sabemos que essa é uma nova variável e que estamos colocando um tipo dentro dela.

````go
func main() {
    contaDoGuilherme := ContaCorrente{}

    fmt.Println(contaDoGuilherme)
}
````

Como não temos nenhum erro, rodaremos o código de novo e ainda vamos ver o mesmo resultado.

Mas não é nosso objetivo. Queremos obter todos os campos. Mas já temos a contaCorrente{}, abrimos e fechamos chaves. Podemos colocar dentro dessas chaves os campos e as informações do cliente.

````go
func main() {
    contaDoGuilherme := ContaCorrente{titular: "Guilherme",
        numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}


    fmt.Println(contaDoGuilherme)
}
````

Vamos salvar e não haverá nenhuma mensagem de erro. Executaremos o código mais uma vez com go run main.go e será impresso "Guilherme 589 123456 125.5", as informações que queríamos.

Se quisermos criar uma outra conta, por exemplo, a conta da Bruna, precisamos declarar contaDaBruna como do tipo contaCorrente{} por meio dos sinais ":= " . Dentro das chaves já não precisaremos passar o nome do campo e o valor do campo. Já podemos passar os valores diretamente, na ordem em que foram criados. A titular será "Bruna", Agência 222, Conta 111222 e Saldo 200. Exibiremos o print da contaDaBruna na última linha do código.

````go
func main() {
    contaDoGuilherme := ContaCorrente{titular: "Guilherme",
        numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

    contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

    fmt.Println(contaDoGuilherme)
    fmt.Println(contaDaBruna)
}
````

Salvaremos, vamos executar e teremos o mesmo resultado, as informações da Bruna de acordo com os campos.

Então, queremos usar sempre o segundo modo, pois com ele não precisaremos escrever o nome de todos os campos novamente. Algo interessante é que quando temos uma estrutura e vamos necessariamente todos os campos, podemos usar o segundo modo, mais sucinto, e será entendido que estamos passando as informações na ordem que passamos os campos a princípio.

Se não precisarmos de todos os campos, por exemplo, apenas titular e saldo, teríamos que escrevê-los para especificar que os valores se tratam dos dois. Na conta do Guilherme, vamos tirar numeroConta e numeroAgencia.

````go
contaDoGuilherme := ContaCorrente{titular: "Guilherme", saldo: 125.5}

fmt.Println(contaDoGuilherme)
````

Vamos rodar e aparecerá no terminal o nome do Titular, Guilherme, no Número da Agência e da Conta terão sido colocados "0"s e o saldo aparecerá também.

Temos essas duas opções de utilização da nossa estrutura. Do outro modo, teríamos 8 variáveis referentes aos campos do Guilherme e da Bruna. Dessa forma, com poucas linhas de código conseguimos um resultado mais elegante usando a linguagem Go.

O código ficará disponível na última atividade da aula atual, em O que aprendemos? para a utilização.

- Código completo:

~~~~go
package main

import (
	"fmt"
)

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoGuilherme := ContaCorrente{titular: "Guilherme",
		numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)
}
~~~~

## RESUMO

Nesta aula, aprendemos como utilizar structs em Go para agrupar diferentes tipos de dados em uma única estrutura.
Vimos que, ao criar uma struct, podemos definir seus campos e seus respectivos tipos. Ao instanciar uma struct, podemos fornecer valores para esses campos de duas maneiras:
Nomeando os campos: Especificamos o nome de cada campo seguido por seu valor. Essa forma é útil quando não queremos fornecer valores para todos os campos ou quando queremos especificar apenas alguns campos.
Sem nomear os campos: Fornecemos os valores na ordem em que os campos foram definidos na struct. Essa forma é mais concisa, mas exige que forneçamos valores para todos os campos.

Além disso, vimos que, quando não fornecemos um valor para um campo, o Go atribui um valor padrão para esse campo, dependendo de seu tipo.

Em resumo, as structs nos permitem organizar e manipular dados de forma mais eficiente em Go, agrupando informações relacionadas em uma única estrutura.

## 📋 **RESUMO SIMPLES**

**Structs** em Go permitem agrupar dados relacionados em uma única estrutura. 

**Duas formas de inicializar:**
- **Com campos nomeados:** `ContaCorrente{titular: "João", saldo: 100.0}`
- **Por ordem:** `ContaCorrente{"João", 589, 123456, 100.0}`

**Zero values:** Campos não preenchidos recebem valores padrão (`""` para string, `0` para int/float).

**Vantagem:** Organiza múltiplas variáveis relacionadas em uma única estrutura, tornando o código mais limpo e legível.