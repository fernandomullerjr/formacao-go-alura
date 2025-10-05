package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 2
const delay = 5

func main() {
	exibeIntroducao()
	// registraLog("teste-site-falso", false) // Chamada de exemplo para a função registraLog // IGNORE
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
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	// Usando a função para ler sites do arquivo ao invés do slice fixo
	sites := leSitesDoArquivo()

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
	// Capturando o segundo parâmetro (erro)
	resp, err := http.Get(site)
	// Verificando se houve algum erro
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		return
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	// Declara um slice vazio de strings para armazenar todos os sites
	// que serão lidos do arquivo sites.txt
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		return sites
	}

	// defer arquivo.Close() agenda o fechamento do arquivo para ser executado
	// automaticamente quando a função terminar, independentemente de como ela termine
	// (return normal, panic, ou qualquer outro tipo de saída da função).
	// Isso garante que o arquivo seja sempre fechado, evitando vazamento de recursos
	// e liberando o handle do arquivo para o sistema operacional.
	defer arquivo.Close()

	// Cria um leitor bufferizado (com buffer) para ler o arquivo de forma eficiente.
	// O bufio.NewReader() permite ler o arquivo linha por linha sem carregar
	// todo o conteúdo na memória de uma vez, o que é útil para arquivos grandes.
	// Ele fornece métodos como ReadString() para leitura linha a linha.
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		// Só adiciona linhas não vazias
		if linha != "" {
			sites = append(sites, linha)
		}

		if err == io.EOF {
			break
		}
	}

	return sites
}

func registraLog(site string, status bool) {

	// Abre (ou cria) o arquivo "log.txt" com as seguintes flags:
	// os.O_CREATE: cria o arquivo se ele não existir
	// os.O_RDWR: permite leitura e escrita no arquivo
	// os.O_APPEND: adiciona conteúdo ao final do arquivo (não sobrescreve)
	// 0666: define permissões de leitura e escrita para owner, group e others
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	//fmt.Println(arquivo)		// <-- Esta linha está imprimindo a referência do arquivo

	// Escreve uma linha no arquivo concatenando com o nome do site, o texto " - online: "
	// o status convertido de bool para string usando strconv.FormatBool()
	// uma quebra de linha "\n" para separar os registros
	// Adicionando timestamp formatado "dd/mm/aaaa hh:mm:ss" no início do log
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site +
		" - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}
