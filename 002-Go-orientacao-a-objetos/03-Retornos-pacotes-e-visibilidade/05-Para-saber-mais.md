05
Para saber mais
Próxima Atividade

Executando um projeto localmente

Quando realizamos o download de um projeto Go, antes de executar o código, um passo importante é necessário:

Exibindo o código main do github destacando o nome do usuário gopath
Alterando nome do caminho das importações

Observe que o nome do caminho da importação é alura. Para executar a aplicação em sua máquina, altere adicionando o seu nome de usuário de sistema.



- The problem is that Go cannot find the local package because the module may not be initialized or the import path doesn't match the module name in go.mod.

go mod init github.com/fernandomullerjr/formacao-go-alura/002-Go-orientacao-a-objetos/03-Retornos-pacotes-e-visibilidade/go_oo-aula3

go mod tidy

> go mod init github.com/fernandomullerjr/formacao-go-alura/002-Go-orientacao-a-objetos/03-Retornos-pacotes-e-visibilidade/go_oo-aula3
go: creating new go.mod: module github.com/fernandomullerjr/formacao-go-alura/002-Go-orientacao-a-objetos/03-Retornos-pacotes-e-visibilidade/go_oo-aula3
go: to add module requirements and sums:
        go mod tidy
> go mod tidy
> date
Sat Nov 22 22:07:01 -03 2025
