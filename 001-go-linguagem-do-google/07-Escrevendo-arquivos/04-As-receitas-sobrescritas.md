# 04 As receitas sobrescritas

João, nosso companheiro do último exercício, viu que apesar de estar conseguindo criar o seu arquivo de receitas.txt com a função abaixo:

os.OpenFile("receitas.txt", os.O_RDWR|os.O_CREATE, 0666)

Toda vez que ele tentava adicionar uma nova receita com a função .WriteString("Nova receita aqui") ele acabava por sobrescrever a receita antiga e nunca conseguia aumentar o seu livro de receitas.

O que o João está fazendo de errado para não conseguir adicionar mais de uma receita no seu arquivo de texto ?
Selecione uma alternativa

    Ele está utilizando a função errada para escrever no arquivo, a função .WriteString() sempre escreve por cima do que já está no arquivo.

    Ele esqueceu de colocar a flag de os.O_APPEND , que iria permitir ele adicionar novos textos no arquivo em vez de sobrescrevê-lo.

    Ele abriu o arquivo com a função os.OpenFile, o que vai impedir que ele adicione novos textos ao arquivo, ele só vai conseguir começar a escrever do começo do arquivo, o que faz com que tudo que ele escreva sobrescreva o que estava antes.


## RESPOSTA
Ele esqueceu de colocar a flag de os.O_APPEND , que iria permitir ele adicionar novos textos no arquivo em vez de sobrescrevê-lo.

Correto! A flag os.O_APPEND é essencial quando queremos adicionar textos em um arquivo, em vez de começar sempre escrevendo do começo do mesmo.
