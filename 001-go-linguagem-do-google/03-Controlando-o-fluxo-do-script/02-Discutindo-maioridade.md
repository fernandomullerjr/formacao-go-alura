# 02 Discutindo maioridade...


Para qual das alternativas abaixo o Go vai permitir compilarmos o if abaixo?

~~~~go
if maiorDeIdade {

}
~~~~

    Alternativa incorreta

    maiorDeIdade := true

Já que a variável é avaliada pelo Go como booleana por causa do valor de inicialização true, o if funcionará.
Alternativa incorreta

var idade int
maiorDeIdade = idade >= 21

Alternativa incorreta

maiorDeIdade := 23

Alternativa incorreta

idade int
maiorDeIdade := idade > 21

Lembre-se que o fluxo condicional if no Go só permite expressões booleanas.
