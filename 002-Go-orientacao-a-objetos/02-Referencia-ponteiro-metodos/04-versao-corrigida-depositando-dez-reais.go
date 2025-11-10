package main

import (
	"fmt"
)

type Conta struct {
	saldo float64
}

func (c *Conta) depositarDezReais() {
	c.saldo += 10
}

func main() {
	contaTeste := Conta{saldo: 10}

	fmt.Println("Saldo inicial:", contaTeste.saldo)

	contaTeste.depositarDezReais()
	fmt.Println("Após primeiro depósito:", contaTeste.saldo)

	contaTeste.depositarDezReais()
	fmt.Println("Após segundo depósito:", contaTeste.saldo)

	fmt.Println("Conta final:", contaTeste)
}
