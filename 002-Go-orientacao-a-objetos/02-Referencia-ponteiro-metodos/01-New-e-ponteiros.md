# 01 New e ponteiros

## Transcrição

Vimos duas formas de utilizar a struct, a primeira passando o campo e o valor que queremos armazenar dentro dessas variáveis e a segunda passando o conteúdo sem especificar os campos, desde que eles estejam na ordem em que foram declarados.

Será que só existem essas formas de utilizar a struct? Não, há a opção de usá-la de forma similar a outras linguagens de programação como Java e C#. Aprenderemos como.

Nosso primeiro passo será declarar uma variável chamada contaDaCris. Ela será do tipo ContaCorrente. Na linha seguinte, faremos contaDaCris = new(). A palavra new é bastante conhecida de quem já programa em Java ou C#. Dentro dos parênteses passaremos o tipo, que é ContaCorrente.

A partir de então, conseguiremos atribuir os valores acrescentando um ponto (.) após contaDaCris e o nome do campo na sequência. Então atribuiremos o valor desse campo.

~~~~go
var contaDaCris ContaCorrente
contaDaCris = new(ContaCorrente)
contaDaCris.titular = "Cris"

fmt.Println(contaDaCris)
~~~~

Teclaremos "Command + J" para abrir o terminal e "Ctrl + L" para limpá-lo. O "Command + S" será necessário para salvar o código. Porém, aparecerá um erro, pois não podemos utilizar nosso tipo ContaCorrente já atribuindo um valor. Isso porque temos uma variável do tipo ContaCorrente, a contaDaCris. Mas o código não entendeu que o tipo da variável é o mesmo que está sendo passado para new(). Precisamos identificar que se tratam do mesmo tipo.

Para conseguir dizer que contaDaCris aponta para uma ContaCorrente, colocaremos um asterisco na frente. Teremos var contaDaCris *ContaCorrente.

Agora não teremos mais nenhum erro. Limparemos o terminal e vamos roda de novo. Veremos no terminal:

&{Cris 0 0 0}

Esse resultado se deve a só termos atribuído o nome, então os outros valores tomaram a forma de zero velho. Atribuiremos agora os demais valores.

Mas por que o "&" e por que temos que apontar para ContaCorrente com asterisco? Nas outras alternativas de uso da struct, o código entendia que que a conta corrente do cliente apontava para um tipo, uma estrutura de conta.

Exemplificaremos o que está acontecendo. Vamos imaginar que um edifício tem uma cobertura super bonita, apertamentos maiores no penúltimo e no último andar, e o térreo e o primeiro andar têm todos os apartamentos do mesmo tamanho.

Para cada apartamento há uma caixa de correio que os identifica. As correspondências que chegam para cada apartamento ficam nas respectivas caixas de correio. Eles possuírem tamanhos diferentes é um ponto importante.

Isso quer dizer que nossas caixas de correio são como os nossos ponteiros ou referências. Cada caixa, independentemente do tamanho do apartamento, apontará para o mesmo lugar se tratando do edifício. Os ponteiros também serão todos iguais. Porém, quando falamos de cada apartamento, os ponteiros apontarão para lugares diferentes do prédio.

Para concluir, as caixas do correio serão nossos ponteiros. O edifício será a memória da nossa aplicação, o local em que podemos armazenar informações. Cada apartamento corresponde a um tipo a ser armazenado.

Então existem tipos menores, os relacionados com os apartamentos do térreo e primeiro andar, onde não precisamos de muito espaço para o armazenamento. O segundo andar será semelhante.

No terceiro andar já teremos um apartamento maior, que permite guardar elementos maiores. O ponteiro, no entanto, será do mesmo tamanho. Essa é apenas uma reflexão simples do que ocorre no nosso desenvolvimento.

Então a contaDaCris precisa apontar para ContaCorrente. O código em Go não entenderá corretamente se tirarmos e asterisco e salvarmos a aplicação, pois ficará sem entender para onde new está apontando, se é a conta corrente da Cris ou uma nova conta. A partir do momento em que o ponteiro é colocado, é como se alocássemos um espaço e disséssemos que a caixa do correio aponta para um apartamento em particular.

A caixa de correio contaDaCris apontará para o apartamento ContaCorrente . Podemos criar cada um dos campos da conta corrente de Cris de acordo com a struct. Criaremos o saldo nesse momento, com contadaCris.saldo = 500.

O último detalhe será o "&" que aparece no terminal quando fazemos a impressão do que temos. Isso porque seguindo o mesmo raciocínio do nosso exemplo, &{Cris 0 0 500} indica que os campos são um conteúdo que está dentro do apartamento. Mas "&" não é interessante para nós, somente o conteúdo.

Por isso, colocaremos o asterisco em contaDaCris também na hora da impressão.

~~~~go
var contaDaCris *ContaCorrente
contaDaCris = new(ContaCorrente)
contaDaCris.titular = "Cris"
contaDaCris.saldo = 500


fmt.Println(*contaDaCris)
~~~~

Se limparmos o terminal e rodarmos de novo, o resultado será igual ao que teríamos usando as outras opções para structs, {Cris 0 0 500} somente.

Então, por debaixo dos panos ocorrerá todo o processo referente aos ponteiros e à alocação da memória nos outros códigos também.

A pergunta que fica é: com qual forma de usar structs é mais fácil trabalhar. Podemos usar a técnica de escrever contaDoGuilherme := ContaCorrente{titular: "Guilherme", saldo: 125.5} quando provavelmente não vamos precisar sempre de todos os campos preenchidos.

A segunda forma, contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200} é mais indicada quando todos os campos serão utilizáveis. Esses serão os modos mais compatíveis com o Go.

Mas o terceiro modo também é importante sabermos usar, pois se precisar lidar com um projeto com esse modelo, saberá do que se tratam os ponteiros e como manipular esses dados.

Essa foi uma aula para termos ideia de que existe essa alternativa e como utilizá-la, mas nos próximos momentos trabalharemos sempre com o primeiro ou o segundo jeitos de lidar com structs. `

- Efetuando teste:

~~~~bash
> go run 01-ponteiros.go
{Guilherme 589 123456 125.5}
{Bruna 222 111222 200}
{Cris 0 0 500}
> date
Sat Oct 25 22:48:21 -03 2025
~~~~

## RESUMO

Nesta aula, o instrutor aborda uma terceira forma de utilizar structs em Go, similar a linguagens como Java e C#, utilizando a palavra-chave new e ponteiros.
Foi demonstrado como declarar uma variável do tipo ContaCorrente e alocar memória para ela usando new, e como atribuir valores aos campos dessa struct através de ponteiros.Para ajudar na compreensão, ele utilizou uma analogia com um edifício e caixas de correio, onde o edifício representa a memória, os apartamentos representam os tipos de dados e as caixas de correio representam os ponteiros.
Além disso, foi explicado o significado do "&" na impressão dos valores e como usar o asterisco para acessar o conteúdo da struct.
Embora essa forma de usar structs seja válida, o instrutor menciona que as formas mais comuns e compatíveis com Go são as que foram apresentadas em aulas anteriores, mas é importante conhecer essa terceira opção caso precise lidar com projetos que a utilizem.

