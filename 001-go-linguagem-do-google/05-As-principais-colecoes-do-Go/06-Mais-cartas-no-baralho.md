
Mais cartas no baralho
Ícone da Luri
Desafio proposto pela Luri, a IA da Alura.
Próxima Atividade

No exercício anterior, Reinaldo criou um slice que representava os pontos utilizados numa técnica de estimativa chamada Planning Poker.

pontosPlanningPoker := [] int {1, 2, 3, 5, 8, 13, 21}

Contudo, ele se esqueceu de colocar uma última pontuação, 40. Em vez de mudar o código de inicialização do slice, usou a função append.

pontosPlanningPoker = append(pontosPlanningPoker, 40)

O que será impresso no terminal quando ele executar o seguinte código:

fmt.Println(cap(pontosPlanningPoker))

14

Você acertou em cheio!

Fernando, sua resposta está correta, parabéns! A função append em Go, quando precisa adicionar um novo elemento a um slice que já atingiu sua capacidade máxima, aloca um novo array com uma capacidade maior. Nesse caso, a capacidade do slice pontosPlanningPoker é aumentada para acomodar o novo elemento 40, e o cap passa a ser 14.
Continue praticando e aprofundando seus conhecimentos em Go, você está no caminho certo!