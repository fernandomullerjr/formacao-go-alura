# 02 Tipos aninhados

## Transcrição

Criamos a struct de Titular no pacote de clientes e indicamos na ContaCorrente do pacote de contas que um Titular pertencerá a um outro tipo que geramos, o tipo Titular.

Agora queremos criar uma ContaCorrente e utilizar as duas estruturas que já criamos. Para isso, removeremos contaDaSilvia e contaDoGustavo que usamos em exemplos anteriores e criaremos uma nova conta em "main.go", contaDoBruno.

Essa conta virá do pacote de "contas" e será do tipo ContaCorrente. O primeiro campo dessa nossa estrutura de conta é o Titular, então escreveremos isso e precisaremos indicar que que ele vem do pacote de "clientes". Mas como faremos isso se não temos o titular?

Será necessário importar também o pacote "clientes" e faremos a indicação Titular: cliente.Titular.

Esse Titular terá um Nome, "Bruno", e um CPF, também do tipo string, que será "123.111.123.12". Por fim, ele também terá uma Profissao, "Desenvolvedor".

Após a primeira estrutura referente a Titular, colocaremos a primeira chave fechando o conteúdo desses campos e uma vírgula. Então, acrescentaremos os campos restantes que serão relacionados à ContaCorrente em vez de Titular.

O NumeroAgencia será 123, NumeroConta 123456 e por fim, o Saldo de 100.

Criaremos um print para visualizar essa conta.

```go
import (
    "fmt"

    "github.com/alura/banco/contas"
    "github.com/alura/banco/clientes"
)

func main() {
    contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{
        Nome: "Bruno",
        CPF: "123.111.123.12",
        Profissao: "Desenvolvedor"},
        NumeroAgencia: 123, NumeroConta: 123456, Saldo: 100}

    fmt.Println(contaDoBruno)
}
```

Vamos salvar e rodar o código com `go run main.go` no terminal. No momento em que executamos, veremos duas chaves. A primeira chave será referente à ContaCorrente e a seguinte ao Titular, porque ele trata de duas estruturas diferentes.

```
{{Bruno 123.111.123.12 Desenvolvedor} 123.123456 100}
```

Agora, se colocarmos o cursor do mouse sobre a ContaCorrente ou do Titular, visualizamos os campos e os títulos.

Para que a compreensão do nosso código seja facilitada podemos criar primeiro o cliente "Bruno" e depois a conta dele.

Faremos o mesmo processo, mas no lugar de criar contaDoBruno, criaremos o clienteBruno que será dois pontos igual (`:=`) `clientes.Titular{}`, sendo Titular o nosso tipo. Teremos os campos de nome, CPF e profissão e passaremos o conteúdo entre as chaves.

Criado o cliente, agora criaremos a conta, contaDoBruno. Ela será do tipo ConteCorrente do pacote de "contas". Dentro das chaves, colocaremos o conteúdo dos campos. clienteBruno, 123, 123456, 100

```go
func main() {
    clienteBruno := clientes.Titular{"Bruno", "123.123.123.12", "Desenvolvedor"}
    contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100}
    fmt.Println(contaDoBruno)
}
```

Vamos salvar e não veremos nenhuma mensagem de erro. Teremos exatamente o mesmo resultado no console:

```
{{Bruno 123.111.123.12 Desenvolvedor} 123.123456 100}
```

Porém, agora nosso código ficou muito mais legível para conseguirmos identificar e criar nossas estruturas.




## Dia 29/11/2025

- Testando a abordagem de leitura mais dificil:

~~~~go
package main

import (
	"fmt"

	"github.com/fernandomullerjr/formacao-go-alura/002-Go-orientacao-a-objetos/04-Composicao-e-encapsulamentos/02-Tipos-aninhados-codigo/clientes"
	"github.com/fernandomullerjr/formacao-go-alura/002-Go-orientacao-a-objetos/04-Composicao-e-encapsulamentos/02-Tipos-aninhados-codigo/contas"
)

func main() {
	contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Bruno",
		CPF:       "123.111.123.12",
		Profissao: "Desenvolvedor"},
		NumeroAgencia: 123, NumeroConta: 123456, Saldo: 100}

	fmt.Println(contaDoBruno)
}
~~~~


- Resultado:

>
> go run main.go
{{Bruno 123.111.123.12 Desenvolvedor} 123 123456 100}
> date
Sat Nov 29 21:49:03 -03 2025



- Testando a abordagem que deixa mais legivel o código, bem melhor:

~~~~go
package main

import (
	"fmt"

	"github.com/fernandomullerjr/formacao-go-alura/002-Go-orientacao-a-objetos/04-Composicao-e-encapsulamentos/02-Tipos-aninhados-codigo/clientes"
	"github.com/fernandomullerjr/formacao-go-alura/002-Go-orientacao-a-objetos/04-Composicao-e-encapsulamentos/02-Tipos-aninhados-codigo/contas"
)

func main() {
	clienteBruno := clientes.Titular{"Bruno", "123.123.123.12", "Desenvolvedor"}
	contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100}
	fmt.Println(contaDoBruno)
}

~~~~

- Resultado:

>
> go run main.go
{{Bruno 123.123.123.12 Desenvolvedor} 123 123456 100}
>
> date
Sat Nov 29 21:50:03 -03 2025