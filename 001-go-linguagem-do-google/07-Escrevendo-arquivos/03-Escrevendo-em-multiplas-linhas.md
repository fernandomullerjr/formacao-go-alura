
# 03 Escrevendo em múltiplas linhas

## Transcrição

Agora que conseguimos criar e abrir o arquivo na função `registraLog`, não podemos nos esquecer de fechá-lo:

````go
// restante do código omitido

func registraLog(site string, status bool) {

    arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR, 0666)

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    fmt.Println(arquivo)

    arquivo.Close()
}
````

Agora, ao invés de imprimir o ponteiro para o arquivo, queremos escrever nele. Vamos começar escrevendo somente se o site carregou ou não.

## Escrevendo no arquivo

Para escrever no arquivo, ele possui a função `WriteString`, ou seja, basta passar o texto para essa função, que ela o escreve dentro do arquivo:

````go
// restante do código omitido

func registraLog(site string, status bool) {

    arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR, 0666)

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    arquivo.WriteString(site + " - online: " + status + "\n")

    arquivo.Close()
}
````

Mas o Go irá reclamar, isso acontece pois a função `WriteString` recebe uma string, e nós estamos tentando concatenar com um booleano, então precisamos convertê-lo com a função strings. Para converter um tipo booleano em uma string, vamos utilizar o pacote `strconv`, que possui a função `FormatBool`:

````go
// restante do código omitido

func registraLog(site string, status bool) {

    arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR, 0666)

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    arquivo.WriteString(site + " - online: " + strconv.FormatBool(status) + "\n")

    arquivo.Close()
}
````

Ao iniciar o monitoramento e ficar verificando o arquivo de logs, vemos que as linhas são sobrescritas a cada site novo monitorado. Isso acontece porque por padrão, quando escrevemos em um arquivo, é sempre escrito do seu começo. Para isso, na hora de abrir o arquivo, existe a flag `O_APPEND`, que faz com que o texto seja escrito ao final do arquivo:

````go
// restante do código omitido

func registraLog(site string, status bool) {

    arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    arquivo.WriteString(site + " - online: " + strconv.FormatBool(status) + "\n")

    arquivo.Close()
}
````

Agora conseguimos imprimir corretamente no arquivo, faltando escrever o horário em que o site estava online ou offline. Faremos isso no próximo vídeo :)



## RESUMO

Nesta aula, aprendemos como escrever em arquivos usando a linguagem Go.
Função registraLog:
A função registraLog é utilizada para registrar informações sobre o status de um site em um arquivo de log.
Utiliza os.OpenFile para abrir o arquivo "log.txt" com as flags O_CREATE (cria o arquivo se não existir), O_RDWR (permissão de leitura e escrita) e O_APPEND (escreve no final do arquivo).
Escreve no arquivo usando arquivo.WriteString, formatando a string com o nome do site e seu status (online ou offline).
Converte o status booleano para string usando strconv.FormatBool.
Fecha o arquivo com arquivo.Close() para garantir que os dados sejam gravados e o arquivo seja liberado.


Escrevendo no Arquivo:
A função WriteString é utilizada para escrever texto no arquivo.
Para concatenar um booleano com uma string, é necessário convertê-lo utilizando a função strconv.FormatBool.


Flag O_APPEND:
A flag O_APPEND é utilizada ao abrir o arquivo para garantir que o texto seja escrito ao final do arquivo, em vez de sobrescrever o conteúdo existente.




