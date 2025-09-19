
# 10 Cronômetro

Fábio é um personal trainer e encomendou um programa para marcar o tempo de seus clientes.

O desenvolvedor está tendo dificuldades para fazer o seguinte programa compilar. Você pode ajudá-lo?

~~~~go
package main

import "fmt"

const aquecimento = 5
const corridaLeve = 10
const corridaForte = 10

func main() {
    fmt.Println("Período de alongamento...")
    time.Sleep(alongamento * time.Minute)

    fmt.Println("Período de aquecimento...")
    time.Sleep(aquecimento * time.Minute)

    fmt.Println("Período de corrida leve...")
    time.Sleep(corridaLeve * time.Minute)

    fmt.Println("Período de corrida forte...")
    time.Sleep(corridaForte * time.Minute)

    fmt.Println("Período de alongamento...")
    time.Sleep(alongamento * time.Minute)
}
~~~~

    Alternativa incorreta

    Falta incluir a constante alongamento:

    const alongamento = 1

Ainda está faltando importar o pacote time para que o programa compile.
Alternativa incorreta

Falta importar o pacote time...

~~~~go
import "time"
~~~~

...e incluir a constante alongamento:

~~~~go
const alongamento = 1
~~~~

Isso aí!
Alternativa incorreta

Falta importar o pacote time...

import "time"

Parabéns, você acertou!
Este conteúdo foi útil para o 