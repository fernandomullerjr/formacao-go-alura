package main

import "fmt"

// ContaCorrente represents a bank account
type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	// Exemplo 6: Comparando conteúdo de ponteiros (dereferenciando)
	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	var contaDaCris2 *ContaCorrente
	contaDaCris2 = new(ContaCorrente)
	contaDaCris2.titular = "Cris"
	contaDaCris2.saldo = 500

	fmt.Println("contaDaCris:", contaDaCris)
	fmt.Println("contaDaCris2:", contaDaCris2)
	fmt.Println("*contaDaCris == *contaDaCris2 (comparando conteúdo):", *contaDaCris == *contaDaCris2) // true - conteúdo igual
}
