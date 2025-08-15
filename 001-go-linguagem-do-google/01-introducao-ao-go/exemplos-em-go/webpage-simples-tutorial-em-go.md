fernandomullerjr: como testar este codigo?

Workspace: ## Plan

1. Create a test file that tests the HTTP handler functionality
2. Use `httptest` package to create a test server and recorder
3. Test the root "/" endpoint to verify it returns the expected HTML content
4. Test the HTTP status code (should be 200 OK)
5. Test the response body contains the expected HTML
6. Use standard Go testing patterns with `testing.T`



Made changes.

fernandomullerjr: como executar e abrir esta pagina? o código está correto? quais comandos uso para tentar abrir a pagina no ambiente local?

Workspace: Collecting workspace informationO código está correto, mas usar a porta 80 exige root no Linux. Para rodar localmente, prefira uma porta alta (ex.: 8080) e trate o erro do servidor.

- Arquivo: webpage-basica1.go (função main.main)

````go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, "<h1>Hello, World!</h1>")
	})

	log.Println("Servindo em http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
````

Como executar e abrir:
- Pelo terminal do VS Code:
  - Rodar direto: `go run ./exemplos-em-go/webpage-basica1.go`
  - Ou buildar e executar:
    - `go build -o webpage ./exemplos-em-go`
    - `./webpage`
- Abrir no navegador: http://localhost:8080/
- Testar via curl: `curl -i http://localhost:8080/`

Se quiser manter a porta 80:
- Rodar com sudo: `sudo go run ./exemplos-em-go/webpage-basica1.go` (não recomendado), ou
- Após build: `sudo setcap 'cap_net_bind_service=+ep' ./webpage` e rodar sem sudo.

Similar code found with 1 license type







fernando@debian10x64:~/cursos/golang/exemplos-em-go$
fernando@debian10x64:~/cursos/golang/exemplos-em-go$ go run webpage-basica2.go
2025/08/10 15:40:54 Servindo em http://localhost:8081/





fernando@debian10x64:~$ netstat -plant | grep 808
(Not all processes could be identified, non-owned process info
 will not be shown, you would have to be root to see it all.)
tcp        0      0 0.0.0.0:8080            0.0.0.0:*               LISTEN      -
tcp        0      0 127.0.0.1:48548         127.0.0.1:8081          ESTABLISHED 10019/firefox-esr
tcp6       0      0 :::8081                 :::*                    LISTEN      11729/webpage-basic
tcp6       0      0 ::1:45738               ::1:8081                TIME_WAIT   -
tcp6       0      0 127.0.0.1:8081          127.0.0.1:48548         ESTABLISHED 11729/webpage-basic
fernando@debian10x64:~$
fernando@debian10x64:~$
fernando@debian10x64:~$
fernando@debian10x64:~$
fernando@debian10x64:~$ curl  http://localhost:8081/
<h1>Hello, World!</h1>fernando@debian10x64:~$
fernando@debian10x64:~$ date
Sun 10 Aug 2025 03:43:33 PM -03
fernando@debian10x64:~$


## RESUMO
- Portas
usar a porta 80 exige root no Linux. Para rodar localmente, prefira uma porta alta (ex.: 8080) e trate o erro do servidor.

- Webpage usando 8081
- Comando para subir a página:
cd ~/cursos/golang/exemplos-em-go
go run webpage-basica2.go