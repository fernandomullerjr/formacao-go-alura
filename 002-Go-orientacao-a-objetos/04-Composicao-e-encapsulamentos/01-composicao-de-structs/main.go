package main

import (
	"fmt"

	"github.com/fernandomullerjr/formacao-go-alura/002-Go-orientacao-a-objetos/04-Composicao-e-encapsulamentos/01-composicao-de-structs/clientes"
	"github.com/fernandomullerjr/formacao-go-alura/002-Go-orientacao-a-objetos/04-Composicao-e-encapsulamentos/01-composicao-de-structs/contas"
)

func main() {
	clienteSilvia := clientes.Titular{Nome: "Silvia", CPF: "123.456.789-00", Profissao: "Desenvolvedora"}
	clienteGustavo := clientes.Titular{Nome: "Gustavo", CPF: "987.654.321-00", Profissao: "Analista"}

	contaDaSilvia := contas.ContaCorrente{Titular: clienteSilvia, Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: clienteGustavo, Saldo: 100}

	status := contaDoGustavo.Tranferir(-200, &contaDaSilvia)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}
