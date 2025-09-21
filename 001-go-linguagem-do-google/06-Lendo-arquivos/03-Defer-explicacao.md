O `defer` é uma palavra-chave em Go que **adia a execução de uma função** até o momento em que a função ao redor dela estiver prestes a retornar.

Em termos simples, você diz: "Go, eu quero que você execute esta linha de código, mas só no último segundo, bem antes de sairmos desta função".

### Para que serve?

Seu principal uso é para **tarefas de limpeza**, garantindo que recursos sejam liberados corretamente. Nos seus exemplos:

1.  **`defer arquivo.Close()`** na função `leSitesDoArquivo()`:
    *   **O que faz:** Garante que o arquivo (`sites.txt`) será fechado, não importa como a função termine.
    *   **Por que é importante:** Se você abrir um arquivo e esquecer de fechá-lo, ele fica "preso" pelo seu programa. O `defer` automatiza o fechamento, tornando o código mais seguro e limpo, pois você coloca a instrução de fechar logo após a de abrir.

2.  **`defer resp.Body.Close()`** na função `testaSite()`:
    *   **O que faz:** Garante que o corpo da resposta HTTP (`resp.Body`) será fechado.
    *   **Por que é importante:** Em requisições HTTP, se você não fechar o corpo da resposta, a conexão de rede pode não ser liberada. Fazer isso repetidamente causa "vazamento de recursos" (resource leak), e seu programa pode eventualmente travar por não conseguir abrir novas conexões.

### Vantagens do `defer`:

*   **Código mais limpo:** A instrução de abrir e a de fechar ficam próximas, melhorando a legibilidade.
*   **Menos erros:** Você não corre o risco de esquecer de fechar um recurso em um `if/else` ou em um retorno antecipado. O `defer` cuida disso.