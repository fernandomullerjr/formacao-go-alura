02
Pacote para acesso a internet
Próxima Atividade

Na construção do projeto de monitorização de sites teremos que efetuar uma comunicação com a Web. Isso é facilmente resolvido por que no Go já temos o seguinte pacote:

    Alternativa incorreta

    "net/http"

"net/http" é pacote mais específico à nossa necessidade. Já que nele temos funções para realizar requisições Get e Post.
Alternativa incorreta

"net/web"
Alternativa incorreta

"http"
Alternativa incorreta

"net"
Alternativa incorreta

"io/ioutil"

A utilização de pacotes providos pela linguagem é extremamente comum no desenvolvimento de Software. Com o Go não é diferente! Existem diversos pacotes para os mais variados propósitos.

Para uma comunicação web com um determinado site é necessário a utilização do pacote "net/http". Que pela própria definição dele tem como objetivo fornecer a implementações de cliente e servidor HTTP.

É importante saber que temos vários subdiretórios dentro do "net". Se quiséssemos fazer um envio de email poderíamos usar o "net/smtp".

Como referência para sabermos os pacotes da linguagem temos: https://golang.org/pkg/
