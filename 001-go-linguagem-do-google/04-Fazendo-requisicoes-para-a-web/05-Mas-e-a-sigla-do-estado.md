# 05 Mas e a sigla do estado?

Manoela se empolgou com seu pequeno programa e previu uma melhoria. Agora ela quer que a função retorne também o estado da cidade.

~~~~go
func devolveCidadeEPopulacao() (string, int, bool, string) {
    return "Vila Sem Nome", 4328, true, "RJ"
}
~~~~

Mas por enquanto não quer utilizá-la no programa. O que ela deve fazer na função main abaixo para que o programa compile sem utilizar o estado?
~~~~go
func main() {
    cidade, populacao, capital := devolveCidadeEPopulacao()
    if capital {
        fmt.Println("A capital ", cidade, "tem", populacao, "habitantes")
    } else {
        fmt.Println("A cidade ", cidade, "tem", populacao, "habitantes")
    }
}
~~~~

Selecione uma alternativa

~~~~go
    cidade, populacao, capital, _ := devolveCidadeEPopulacao()
~~~~

~~~~go
    cidade, populacao, capital, null := devolveCidadeEPopulacao()
~~~~
    Nada. Manoela vai precisar necessariamente declarar uma variável para armazenar o estado, e em seguida utilizá-lo no programa. Assim:
~~~~go
    cidade, populacao, capital, estado := devolveCidadeEPopulacao()
    fmt.Println("Estado é", estado)
~~~~



---------------------------------------------------
---------------------------------------------------
---------------------------------------------------
## resposta
R:
~~~~go
    cidade, populacao, capital, _ := devolveCidadeEPopulacao()
~~~~
Quando queremos ignorar um valor retornado em uma função que retorna valores múltiplos, usamos o símbolo _ (underscore).