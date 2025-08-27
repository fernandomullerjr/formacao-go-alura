package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	//exibeNomes()
	//exibeIntroducao()
	for {
		exibeMenu()
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

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := []string{"https://random-status-code.herokuapp.com/",
		"https://www.alura.com.br", "https://www.caelum.com.br"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(site string) {

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
