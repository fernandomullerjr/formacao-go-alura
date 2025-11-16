package main

import (
	"fmt"
)

// ContaCorrente representa uma conta bancária simples.
// Os campos estão em minúsculas, portanto são privados ao pacote atual.
type ContaCorrente struct {
	titular       string  // nome do titular da conta
	numeroAgencia int     // número da agência bancária
	numeroConta   int     // número da conta
	saldo         float64 // saldo disponível
}

// Método "Sacar" — responsável por diminuir o saldo, caso haja valor suficiente.
//
// Receiver: *ContaCorrente (ponteiro)
// → Permite alterar o valor original de 'saldo' dentro da struct.
//
// Parâmetro: valorDoSaque (float64)
// → Quantia que o cliente deseja sacar.
//
// Retorna: string — mensagem de sucesso ou erro.
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	// Regra: só permite saque se o valor for positivo e <= saldo atual
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		// Atualiza o saldo (modifica o valor original, pois c é um ponteiro)
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do depósito menor que zero", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func main() {
	contaDaSilvia := ContaCorrente{titular: "Silvia", saldo: 300}
	contaDoGustavo := ContaCorrente{titular: "Gustavo", saldo: 100}

	status := contaDaSilvia.Transferir(200, &contaDoGustavo)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}
