package lendoarquivos
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	// Lê os sites do arquivo uma vez, no início.
	sites := leSitesDoArquivo()

	// Exibe os sites que serão monitorados.
	fmt.Println("Estes são os sites que serão monitorados:")
	for _, site := range sites {
		fmt.Println("- ", site)
	}
	fmt.Println("") // Adiciona uma linha em branco para formatação

	for {
		exibeMenu() // Descomentado para exibir o menu
		// A chamada foi removida daqui para evitar leituras repetidas.
		comando := leComando()

		switch comando {
		case 1:
			// Passa a lista de sites já lida para a função.
			iniciarMonitoramento(sites)
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

func iniciarMonitoramento(sites []string) {
	fmt.Println("Monitorando...")

	// Não é mais necessário ler o arquivo aqui, pois a lista já foi recebida.
	// sites := leSitesDoArquivo()

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

// func testaSite(site string) {

// 	resp, _ := http.Get(site)

// 	if resp.StatusCode == 200 {
// 		fmt.Println("Site:", site, "foi carregado com sucesso!")
// 	} else {
// 		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
// 	}
// }

func testaSite(site string) {
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(site)
	if err != nil {
		fmt.Println("Erro ao acessar:", site, "-", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	// Tenta abrir o arquivo "sites.txt" para leitura
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		// Se ocorrer erro ao abrir, exibe mensagem e retorna slice vazio
		fmt.Println("Ocorreu um erro:", err)
		return sites
	}
	defer arquivo.Close() // Garante que o arquivo será fechado ao final da função

	// Cria um leitor para ler o arquivo linha por linha
	leitor := bufio.NewReader(arquivo)

	for {
		// Lê uma linha do arquivo até encontrar '\n'
		linha, err := leitor.ReadString('\n')
		// Remove espaços em branco e quebras de linha do final da string
		linha = strings.TrimSpace(linha)
		// Adiciona a linha (site) ao slice de sites
		sites = append(sites, linha)
		// Se chegou ao final do arquivo, encerra o loop
		if err == io.EOF {
			break
		}
	}

    // Fecha o arquivo após terminar a leitura
    arquivo.Close()

	// Retorna o slice com todos os sites lidos do arquivo
	return sites
}