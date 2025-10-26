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
	// Exemplo 2: Comparando structs com conteúdo diferente
	contaDoGuilherme := ContaCorrente{titular: "Guilherme", numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}
	contaDoGuilherme2 := ContaCorrente{titular: "Guilherme", numeroAgencia: 580, numeroConta: 123456, saldo: 125.5}

	fmt.Println("contaDoGuilherme:", contaDoGuilherme)
	fmt.Println("contaDoGuilherme2:", contaDoGuilherme2)
	fmt.Println("contaDoGuilherme == contaDoGuilherme2:", contaDoGuilherme == contaDoGuilherme2) // false
}
