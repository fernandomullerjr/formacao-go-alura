

# 05 Preparando o ambiente: Criando o Workspace

## Criando o seu workspace

Para utilizarmos a linguagem de programação Go devemos seguir algumas convenções, e a principal delas é a sua organização de pastas no seu Workspace.

O Workspace padrão do Go é um diretório chamado /go que fica na pasta do seu usuário em seu sistema operacional. No Windows esta pasta normalmente se encontra em C:/Users/seu-usuario/ , e nos sistemas Unix ( MacOS e distribuições do Linux) normalmente se encontra em /home/seu-usuario/.

Dentro deste diretório devem conter as seguintes pastas:

pasta-do-usuario/
└── go
    ├── bin
    ├── pkg
    └── src

A bin , onde ficará os compilados do nosso código Go.

A pkg, onde ficará os pacotes compartilhados das nossas aplicações, pois o Go é uma linguagem bastante modular, dependendo bastante de pacotes que vamos importando ao longo de nossos códigos.

A src, onde escreveremos o código fonte de cada aplicação.

Crie estas pastas como indicado acima, para que o seu Go Workspace fique funcionando corretamente.