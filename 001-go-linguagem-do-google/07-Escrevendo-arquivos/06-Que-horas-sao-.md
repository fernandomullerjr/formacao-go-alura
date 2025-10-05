06
Que horas são ?
Próxima Atividade

Qual será o tempo exibido no console quando executarmos o código abaixo?

package main

func main(){
    fmt.Println(time.Now().Format("02/01/2006 15:04:05"))
}

Selecione uma alternativa

    Será exibido no horário e data atual, no formato "Dia/Mês/Ano Hora:Minuto:Segundo"

    Será exibido no horário e data atual, no formato ISO.

    Será exibido no horário e data informado no formato ISO, ou seja "2006-01-02T15:04:05Z".


## RESPOSTA
Será exibido no horário e data atual, no formato "Dia/Mês/Ano Hora:Minuto:Segundo"

Exato! O Go tem este estilo diferente de formatar o tempo, que pode ser consultado na sua documentação aqui
https://golang.org/src/time/format.go