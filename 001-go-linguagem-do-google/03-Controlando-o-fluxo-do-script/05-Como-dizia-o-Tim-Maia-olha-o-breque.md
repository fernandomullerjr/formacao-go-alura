# 05 Como dizia Tim Maia, olha o breque!

Em outras linguagens, após cada alternativa é necessário colocar a palavra reservada break. Quais são os comportamentos do compilador da linguagem Go?

    Alternativa correta

    Não precisa colocar break.

No Go no uso do break não é obrigatório.
Alternativa correta

Se o desenvolvedor colocar break, o compilador não vai reclamar.

No Go no uso do break não é obrigatório.
Alternativa correta

O compilador obriga a colocar o break no final de cada alternativa.

Se o desenvolvedor colocar o break, o compilador não vai reclamar. Por exemplo na listagem abaixo.

~~~~go
switch comando {
case 1:
    fmt.Println("Monitorando... ")
    break
case 2:
    fmt.Println("Exibindo logs...")
    break
case 0:
    fmt.Println("Saindo do programa")
    break
default:
    fmt.Println("Não conheço esse comando")
    break
}
~~~~

Porém, ele não é obrigatório. Apenas uma alternativa é executada por avaliação switch.



------------------------------------------------------------------------------
------------------------------------------------------------------------------
------------------------------------------------------------------------------
------------------------------------------------------------------------------
------------------------------------------------------------------------------
------------------------------------------------------------------------------ 
## RESPOSTA
Não precisa colocar break.

No Go no uso do break não é obrigatório.
Alternativa correta

Se o desenvolvedor colocar break, o compilador não vai reclamar.

No Go no uso do break não é obrigatório.
