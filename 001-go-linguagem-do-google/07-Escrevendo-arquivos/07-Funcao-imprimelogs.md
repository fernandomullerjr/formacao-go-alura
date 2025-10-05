07 Função imprimeLogs

Transcrição

Para imprimir os logs do nosso sistema, vamos criar a função imprimeLogs:

// restante do código omitido

func imprimeLogs() {
    
}

Nela, não há nenhuma novidade, precisamos abrir o arquivo de log e imprimir o seu conteúdo. Na aula anterior, vimos como exibir o conteúdo de um arquivo através da função ReadFile, do pacote os, então é ela que usaremos:

// restante do código omitido

func imprimeLogs() {

    arquivo, err := os.ReadFile("log.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    fmt.Println(arquivo)
}

Como a função nos retorna o arquivo em bytes, precisamos convertê-los para string:

// restante do código omitido

func imprimeLogs() {

    arquivo, err := os.ReadFile("log.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    fmt.Println(string(arquivo))
}

O detalhe é que não precisamos fechar o arquivo, pois o pacote os abre o arquivo, o lê e o fecha para nós. Agora, chamamos essa função quando o comando 2 for digitado:

// restante do código omitido

func main() {
    exibeIntroducao()
    for {
        exibeMenu()
        comando := leComando()

        switch comando {
        case 1:
            iniciarMonitoramento()
        case 2:
            fmt.Println("Exibindo Logs...")
            imprimeLogs()
        case 0:
            fmt.Println("Saindo do programa")
            os.Exit(0)
        default:
            fmt.Println("Não conheço este comando")
            os.Exit(-1)
        }
    }
}

Ao executar o programa e digitar o comando 2, os logs são exibidos no terminal!