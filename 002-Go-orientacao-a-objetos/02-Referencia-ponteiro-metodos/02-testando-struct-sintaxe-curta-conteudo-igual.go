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
	// Exemplo 3: Usando sintaxe curta para criação de structs - valores iguais
	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}
	contaDaBruna2 := ContaCorrente{"Bruna", 222, 111222, 200}

	fmt.Println("contaDaBruna:", contaDaBruna)
	fmt.Println("contaDaBruna2:", contaDaBruna2)
	fmt.Println("contaDaBruna == contaDaBruna2:", contaDaBruna == contaDaBruna2) // true
}
