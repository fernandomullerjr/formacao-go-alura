package main

import (
	"fmt"

	"github.com/fernandomullerjr/formacao-go-alura/002-Go-orientacao-a-objetos/04-Composicao-e-encapsulamentos/02-Tipos-aninhados-codigo/clientes"
	"github.com/fernandomullerjr/formacao-go-alura/002-Go-orientacao-a-objetos/04-Composicao-e-encapsulamentos/02-Tipos-aninhados-codigo/contas"
)

func main() {
	clienteBruno := clientes.Titular{"Bruno", "123.123.123.12", "Desenvolvedor"}
	contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100}
	fmt.Println(contaDoBruno)
}
