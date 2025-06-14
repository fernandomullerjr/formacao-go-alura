

## 04 Preparando o ambiente: Instalando o Go

Para instalarmos o Go em nosso computador primeiramente precisamos fazer o Download do binário da linguagem adequado para cada sistema operacional. Visite a página de download oficial da linguagem e faça o download adequado para o seu sistema. Neste curso usaremos a versão 1.8.X do Go:

Imagem dos downloads

Agora vamos seguir para a instalação para cada uma das plataformas:
Instalando para o Windows

Abra o arquivo MSI que contêm a instalação e siga as instruções na tela para instalar o Go. O instalador é bem simples e basta apenas confirmar a instalação com todas as opções que vierem por padrão. O instalador coloca a instalação do Go em c:\Go e não precisamos interferir nisto.

O instalador deve colocar o diretório c:\Go\bin na variável de ambiente PATH do seu sistema. É necessário reiniciar qualquer prompt de comando aberto para que a alteração da variável entre em vigor e você possa testar a instalação.

Ao final da instalação , abra um novo prompt de comando e execute o comando go version e veja se a versão e os comandos do Go foram exibidos corretamente.
Instalando para o MacOS

Baixe o arquivo .pkg, abra-o, e siga as instruções para instalar o Go. Assim como no Windows, o instalador para MacOS é bem simples e apenas devemos confirmar as opções que vierem por padrão e seguir com a instalação. O pacote instala a linguagem Go em /usr/local/go.

O instalador normalmente coloca o diretório /usr/local/go bin na sua variável de ambiente PATH. Talvez seja necessário reiniciar quaisquer sessões abertas do Terminal para que a alteração entre em vigor.

Ao final da instalação , abra um novo prompt de comando e execute o comando go version e veja se a versão e os comandos do Go foram exibidos corretamente.
Instalando para o Ubuntu

Faça o download do arquivo .tar do Go e navegue com o terminal até a pasta aonde o arquivo foi baixado. Use o comando abaixo para extraí-lo em sua pasta /usr/local:

sudo tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz

Repare que neste exemplo estamos utilizando a versão 1.8.3 da linguagem, caso você faça o Download de uma versão mais nova você deve alterar o comando para que o nome do arquivo corresponda ao que foi baixado.

Em seguida precisamos adicionar o caminho /usr/local/go/bin no PATH do sistema. Você pode fazer isto adicionando uma linha extra no arquivo /etc/profile, por exemplo com o editor de textos gedit. Abra o arquivo /etc/profile com o gedit utilizando o comando abaixo:

sudo gedit /etc/profile

E em seguida adicione a seguinte linha no final do arquivo:

export PATH=$PATH:/usr/local/go/bin

Salve a alteração e feche o editor.

Agora utilize o comando:

source /etc/profile

E o Go já deve estar funcionando com sucesso! Faça o teste em seu terminal com o comando go version e veja se a versão e os comandos do Go foram exibidos corretamente.