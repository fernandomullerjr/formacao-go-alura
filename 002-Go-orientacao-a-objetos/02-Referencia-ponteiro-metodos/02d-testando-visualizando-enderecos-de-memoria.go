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
	// Exemplo 7: Visualizando endereços de memória
	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	var contaDaCris2 *ContaCorrente
	contaDaCris2 = new(ContaCorrente)
	contaDaCris2.titular = "Cris"
	contaDaCris2.saldo = 500

	fmt.Printf("Endereço onde contaDaCris aponta: %p\n", contaDaCris)
	fmt.Printf("Endereço onde contaDaCris2 aponta: %p\n", contaDaCris2)
	fmt.Printf("Endereço da variável contaDaCris: %p\n", &contaDaCris)
	fmt.Printf("Endereço da variável contaDaCris2: %p\n", &contaDaCris2)

	fmt.Println("\nComparações:")
	fmt.Println("contaDaCris == contaDaCris2:", contaDaCris == contaDaCris2)     // false - endereços diferentes
	fmt.Println("*contaDaCris == *contaDaCris2:", *contaDaCris == *contaDaCris2) // true - conteúdo igual
	fmt.Println("&contaDaCris == &contaDaCris2:", &contaDaCris == &contaDaCris2) // false - variáveis em locais diferentes
}
