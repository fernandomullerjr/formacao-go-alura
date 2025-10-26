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
	fmt.Println("=== Comparing struct values ===")

	// Example 1: Comparing struct values with same content
	contaDoGuilherme := ContaCorrente{titular: "Guilherme", numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}
	contaDoGuilherme2 := ContaCorrente{titular: "Guilherme", numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	fmt.Println("contaDoGuilherme == contaDoGuilherme2:", contaDoGuilherme == contaDoGuilherme2) // true

	// Example 2: Comparing struct values with different content
	contaDoGuilherme3 := ContaCorrente{titular: "Guilherme", numeroAgencia: 580, numeroConta: 123456, saldo: 125.5}
	fmt.Println("contaDoGuilherme == contaDoGuilherme3:", contaDoGuilherme == contaDoGuilherme3) // false

	fmt.Println("\n=== Comparing struct values (short syntax) ===")

	// Example 3: Using short syntax for struct creation
	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}
	contaDaBruna2 := ContaCorrente{"Bruna", 222, 111222, 200}

	fmt.Println("contaDaBruna == contaDaBruna2:", contaDaBruna == contaDaBruna2) // true

	// Example 4: Different values
	contaDaBruna3 := ContaCorrente{"Bruna", 111, 111222, 200}
	fmt.Println("contaDaBruna == contaDaBruna3:", contaDaBruna == contaDaBruna3) // false

	fmt.Println("\n=== Comparing pointers ===")

	// Example 5: Creating pointers with new()
	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	var contaDaCris2 *ContaCorrente
	contaDaCris2 = new(ContaCorrente)
	contaDaCris2.titular = "Cris"
	contaDaCris2.saldo = 500

	// Display pointer content
	fmt.Println("contaDaCris content:", contaDaCris)   // &{Cris 0 0 500}
	fmt.Println("contaDaCris2 content:", contaDaCris2) // &{Cris 0 0 500}

	// Compare pointer addresses (different)
	fmt.Println("contaDaCris == contaDaCris2:", contaDaCris == contaDaCris2) // false

	// Compare pointer content (same)
	fmt.Println("*contaDaCris == *contaDaCris2:", *contaDaCris == *contaDaCris2) // true

	fmt.Println("\n=== Memory addresses ===")

	// Display memory addresses
	fmt.Printf("Address of contaDaCris: %p\n", contaDaCris)
	fmt.Printf("Address of contaDaCris2: %p\n", contaDaCris2)
	fmt.Printf("Address of contaDaCris variable: %p\n", &contaDaCris)
	fmt.Printf("Address of contaDaCris2 variable: %p\n", &contaDaCris2)

	fmt.Println("\n=== Summary ===")
	fmt.Println("- Struct values: comparison checks content")
	fmt.Println("- Pointers: comparison checks memory addresses")
	fmt.Println("- Dereferenced pointers (*pointer): comparison checks content")
}
