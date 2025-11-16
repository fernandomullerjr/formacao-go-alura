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

func main() {
	// Cria uma instância (objeto) de ContaCorrente com zero values.
	contaDaSilvia := ContaCorrente{}

	// Define manualmente alguns campos.
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	// Mostra o saldo inicial.
	fmt.Println(contaDaSilvia.saldo) // 500

	// Tenta sacar 400 reais — o método altera o saldo original.
	fmt.Println(contaDaSilvia.Sacar(400)) // "Saque realizado com sucesso"

	// Exibe o novo saldo após o saque.
	fmt.Println(contaDaSilvia.saldo) // 100
}
