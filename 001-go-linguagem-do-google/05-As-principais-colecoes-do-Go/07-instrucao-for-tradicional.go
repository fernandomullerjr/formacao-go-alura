package asprincipaiscolecoesdogo
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeNomes()
	//exibeIntroducao()
	for {
		//exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Douglas"
	versao := 1.2
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	// Criando um slice de strings com URLs dos sites a serem monitorados
	// Slice permite adicionar/remover sites dinamicamente se necessário
	sites := []string{
		"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br",
		"https://www.caelum.com.br",
	}

	// FOR TRADICIONAL: Percorre o slice usando índice
	// i = 0: primeira posição do slice
	// i < len(sites): condição para não ultrapassar o tamanho do slice
	// i++: incrementa o índice a cada iteração
	for i := 0; i < len(sites); i++ {
		// Acessa o elemento na posição i do slice
		fmt.Println(sites[i])
	}

	// CÓDIGO LEGADO: Verifica apenas um site específico (será refatorado)
	// TODO: Mover esta lógica para dentro do loop para monitorar todos os sites
	site := "https://www.alura.com.br"

	// http.Get() faz uma requisição HTTP GET para o site
	// O segundo valor de retorno (erro) está sendo ignorado com "_"
	resp, _ := http.Get(site)

	// Verifica o código de status HTTP da resposta
	// 200 = OK (sucesso), outros códigos indicam problemas
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

func devolveNomeEIdade() (string, int) {
	nome := "Douglas"
	idade := 24
	return nome, idade
}

func exibeNomes() {
	// Criando um slice de strings com 3 elementos iniciais
	// Slice é uma estrutura dinâmica, diferente de array que tem tamanho fixo
	nomes := []string{"Douglas", "Daniel", "Bernardo"}

	// len() retorna o número atual de elementos no slice
	fmt.Println("O meu slice tem", len(nomes), "itens")

	// cap() retorna a capacidade atual do slice (espaço alocado na memória)
	// Inicialmente, capacidade = comprimento = 3
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	// append() adiciona um novo elemento ao slice
	// Se necessário, Go automaticamente aumenta a capacidade do slice
	nomes = append(nomes, "Aparecida")

	// Agora temos 4 elementos (len = 4)
	fmt.Println("O meu slice tem", len(nomes), "itens")

	// A capacidade dobrou automaticamente para 6 (Go otimiza o crescimento)
	// Isso evita realocações frequentes de memória
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
}
