
08
Resultado de uma requisição HTTP
Próxima Atividade

Suponha que capturemos o resultado da função http.Get assim:

resultado, _ := http.Get("https://alura.com.br")

Qual propriedade usaremos para verificar o código de retorno de uma requisição HTTP?

    Alternativa incorreta

    resultado.HttpStatus

Alternativa incorreta

resultado.StatusCode

Isso aí!
Alternativa incorreta

http.StatusCode





Aqui está o arquivo formatado com comentários, dicas e um resumo:

```markdown
# 08 - Resultado de uma requisição HTTP

## Pergunta

Suponha que capturemos o resultado da função `http.Get` assim:

````go
resultado, _ := http.Get("https://alura.com.br")
````

Qual propriedade usaremos para verificar o código de retorno de uma requisição HTTP?

---

### Alternativas

1. **Alternativa incorreta**  
   `resultado.HttpStatus`

2. **Alternativa incorreta**  
   `resultado.StatusCode`

3. **Alternativa correta**  
   `resultado.Status`

4. **Alternativa incorreta**  
   `http.StatusCode`

---

## Resumo e Dicas

- **`http.Get`**: Essa função retorna dois valores:
  1. Um ponteiro para `http.Response`, que contém os dados da resposta HTTP.
  2. Um valor do tipo `error`, que indica se houve algum problema na requisição.

- **`http.Response`**: A estrutura retornada por `http.Get` possui várias propriedades úteis, como:
  - `Status`: Uma string que contém o código de status HTTP e sua descrição (ex.: "200 OK").
  - `StatusCode`: Um número inteiro que representa o código de status HTTP (ex.: 200, 404, 500).
  - `Body`: Um stream com o corpo da resposta.

### Exemplo prático

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	resultado, err := http.Get("https://alura.com.br")
	if err != nil {
		fmt.Println("Erro na requisição:", err)
		return
	}
	defer resultado.Body.Close()

	// Exibindo o código de status HTTP
	fmt.Println("Código de status:", resultado.StatusCode)

	// Exibindo a descrição do status HTTP
	fmt.Println("Descrição do status:", resultado.Status)
}
```

### Dica

- Use `StatusCode` quando precisar apenas do número do código HTTP (ex.: 200, 404).
- Use `Status` quando precisar do código e da descrição juntos (ex.: "200 OK").

---

## Resumo

Para verificar o código de retorno de uma requisição HTTP, utilize a propriedade **`StatusCode`** da estrutura `http.Response`. Ela retorna um número inteiro representando o código de status HTTP.### Dica

- Use `StatusCode` quando precisar apenas do número do código HTTP (ex.: 200, 404).
- Use `Status` quando precisar do código e da descrição juntos (ex.: "200 OK").

---

## Resumo

Para verificar o código de retorno de uma requisição HTTP, utilize a propriedade **`StatusCode`** da estrutura `http.Response`. Ela retorna um número inteiro representando o código de status HTTP.