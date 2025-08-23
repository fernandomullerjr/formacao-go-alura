# 06 Para saber mais: Status HTTP 

    A partir desta aula, o curso utiliza a API (https://random-status-code.herokuapp.com/) para verificação de status code que, neste momento, não está mais no ar. Veja abaixo como substituir a API usada no vídeo em suas atividades:

Uma opção alternativa que você pode usar é o seguinte serviço:

    Httpbin

Por meio do Httpbin, é possível simular o status HTTP ao passar o código de status correspondente:

    https://httpbin.org/status/200
    https://httpbin.org/status/404

No código, você pode alternar o status code para verificar a mudança no console:

~~~~go
func iniciarMonitoramento() {
    fmt.Println("Monitorando...")
    // site com URL inexistente
    site := "https://httpbin.org/status/404" // ou 200
    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}
~~~~

Realizando essas implementações é possível atingir os mesmos objetivos que na aula.


- Testando

~~~~bash
fernando@debian10x64:~/cursos/golang/formacao-go-alura/001-go-linguagem-do-google/04-Fazendo-requisicoes-para-a-web$ go run 06-monitoramento-http.go
Olá, sr. Douglas
Este programa está na versão 1.2
1- Iniciar Monitoramento
2- Exibir Logs
0- Sair do Programa
Douglas tem 24 anos
1
O comando escolhido foi 1
Monitorando...
Site: https://httpbin.org/status/404 está com problemas. Status Code: 404
fernando@debian10x64:~/cursos/golang/formacao-go-alura/001-go-linguagem-do-google/04-Fazendo-requisicoes-para-a-web$

fernando@debian10x64:~/cursos/golang/formacao-go-alura/001-go-linguagem-do-google/04-Fazendo-requisicoes-para-a-web$ date
Sat 23 Aug 2025 02:55:12 PM -03
fernando@debian10x64:~/cursos/golang/formacao-go-alura/001-go-linguagem-do-google/04-Fazendo-requisicoes-para-a-web$
~~~~



---------------------------------------------------
---------------------------------------------------
---------------------------------------------------
## RESUMO