# 02 O livro de receitas

João está montando um livro de receitas , e para isto está utilizando sua linguagem favorita, o Go. Para isto, ele quer conseguir escrever no arquivo receitas.txt, que nem sempre vai existir no computador de quem rodar o script do João.

Qual das funções abaixo permite a ele a possibilidade de criar o arquivo receitas.txt caso ele não exista, ou escrever nele caso ele já esteja presente ?
Selecione uma alternativa

    os.Open("receitas.txt", os.O_RDWR|os.O_CREATE, 0666)

    os.OpenFile("receitas.txt", O_RDWR|O_CREATE, 0666)

    os.Open("receitas.txt", os.O_RDWR|os.O_CREATE)

    os.OpenFile("receitas.txt", os.O_RDWR|os.O_CREATE, 0666)


## RESPOSTA
os.OpenFile("receitas.txt", os.O_RDWR|os.O_CREATE, 0666)

Correto! A função os.OpenFile é mais poderosa que a função os.Open e nos permite configurar determinadas flags para configurar como o arquivo será manipulado.
