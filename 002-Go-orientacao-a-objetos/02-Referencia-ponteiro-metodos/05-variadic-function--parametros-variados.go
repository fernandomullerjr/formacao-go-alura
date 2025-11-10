package main

import (
	"fmt"
)

func Somando(numeros ...int) int {
	resultadoDaSoma := 0
	for _, numero := range numeros {
		resultadoDaSoma += numero
	}
	return resultadoDaSoma
}

func main() {
	fmt.Println(Somando(1))
	fmt.Println(Somando(1, 1))
	fmt.Println(Somando(1, 1, 1))
	fmt.Println(Somando(1, 1, 2, 4))
}
