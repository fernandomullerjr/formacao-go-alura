# 07 Leitura da entrada do usuário
 

No último vídeo, para capturar o que o usuário escrever, foi visto que existe a função Scanf, do pacote fmt. Ela recebe como argumento um modificador e um ponteiro para a variável que guardará a entrada do usuário, por exemplo:

var idade int
fmt.Scanf("%d", &idade)

A variável idade é do tipo inteiro, logo, só pode receber números inteiros. Então, como o Go sabe que a variável só pode receber números inteiros, o modificador %d, que representa um número inteiro, deixa de ser necessário. Por isso há uma outra função, que não recebe como argumento o modificador.

Essa função é a:

    Alternativa incorreta

    var idade int
    fmt.Read(&idade)

Alternativa incorreta

var idade int
fmt.Check(&idade)

Alternativa incorreta

~~~~go
var idade int
fmt.Scan(&idade)
~~~~

Alternativa correta! A função é a Scan, também do pacote fmt. Ela consegue inferir o tipo digitado pelo usuário, baseado no tipo da variável recebida.
Alternativa incorreta

var idade int
fmt.Sweep(&idade)

Parabéns, você acertou! 