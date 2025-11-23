package main

import (
	"fmt"

	"github.com/fernandomullerjr/formacao-go-alura/002-Go-orientacao-a-objetos/03-Retornos-pacotes-e-visibilidade/go_oo-aula3/contas"
)

func main() {
	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaDoGustavo.Tranferir(-200, &contaDaSilvia)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}
