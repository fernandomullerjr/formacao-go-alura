# 02 Seu Valdomiro foi à feira

Seu Valdomiro foi à feira e comprou 3 frutas. Sua bolsa tinha capacidade para guardar 4 frutas.

~~~~go
var frutas [4] string
frutas[0] = "Abacaxi"
frutas[1] = "Laranja"
frutas[2] = "Morango"
~~~~

O que vai acontecer se executarmos o código abaixo?

~~~~go
fmt.Println(frutas[3])
~~~~

Selecione uma alternativa

    Será impresso uma string vazia.

    O programa não vai compilar porque a "variável" alocada para a última posição do array não foi usada.

    Acontecerá um erro na execução.




## RESPOSTA
Será impresso uma string vazia.

Isso aí!

Quando os arrays são criados, eles assumem os valores padrão para os tipos de seus elementos. No caso, o tipo do array frutas é string e o valor padrão para cada posição do array será vazio. Portanto, o valor impresso será uma string vazia.