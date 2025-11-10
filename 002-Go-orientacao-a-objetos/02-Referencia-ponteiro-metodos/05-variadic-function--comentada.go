package main

import (
	"fmt"
)

// Função variádica que aceita um número variável de parâmetros int
// O parâmetro "numeros ...int" permite passar 0, 1 ou mais argumentos
func Somando(numeros ...int) int {
	// Inicializa a variável que armazenará o resultado da soma
	resultadoDaSoma := 0

	// Itera sobre todos os números recebidos como parâmetro
	// O "range" percorre o slice de números
	// "_" ignora o índice, pois só precisamos do valor
	for _, numero := range numeros {
		// Soma cada número ao resultado acumulado
		resultadoDaSoma += numero
	}

	// Retorna a soma total de todos os números
	// Os sinais += somam o valor à variável e atualizam ela
	return resultadoDaSoma
}

func main() {
	// Exemplos de uso da função variádica com diferentes quantidades de argumentos

	fmt.Println(Somando(1))          // Resultado: 1 (um parâmetro)
	fmt.Println(Somando(1, 1))       // Resultado: 2 (dois parâmetros)
	fmt.Println(Somando(1, 1, 1))    // Resultado: 3 (três parâmetros)
	fmt.Println(Somando(1, 1, 2, 4)) // Resultado: 8 (quatro parâmetros)

	// Também poderíamos chamar sem parâmetros:
	// fmt.Println(Somando()) // Resultado: 0
}
