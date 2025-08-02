package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome = "Douglas"
	var idade = 24
	var versao = 1.1
	fmt.Println("Olá sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("O tipo da variável idade é", reflect.TypeOf(versao))
}
