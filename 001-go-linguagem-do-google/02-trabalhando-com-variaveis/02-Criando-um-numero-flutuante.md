
# 02 Criando um número flutuante


George é dono de uma padaria e criou um programa em Go para exibir os produtos vendidos nela:

~~~~go
package main

import "fmt"

func main() {
    var precoLeite float = 2.99
    var precoOvo float = 3.99
    var precoPao float = 0.99

    fmt.Println("O preço do leite é R$", precoLeite)
    fmt.Println("O preço do ovo é R$", precoOvo)
    fmt.Println("O preço do pão é R$", precoPao)
}
~~~~

Mas ao executar o programa, ele recebe o seguinte erro:

# command-line-arguments
./hello.go:6: undefined: float

O que aconteceu no programa do George para gerar esse erro?

    Alternativa incorreta

    O erro acontece pois o tipo de variável float no Go não existe, e sim suas versões, 32float para números de até 32 bits, e 64float, para números de até 64 bits.

Alternativa incorreta

O erro acontece pois o tipo de variável float no Go não existe, e sim suas versões, float32 para números de até 32 bits, e float64, para números de até 64 bits.

    **Alternativa correta!** O tipo de variável float no Go não existe, porque ele possui tamanhos, de 32 ou 64 bits, que devem ser especificados no momento da sua declaração. Se a variável for de 32 bits, o seu tipo será float32, mas se for de 64 bits, o seu tipo será float64.
    Alternativa incorreta

O erro acontece pois o tipo de variável float no Go não existe, o correto é double.
Alternativa incorreta

O erro acontece pois George esqueceu de colocar ponto e vírgula ao final das linhas de código.