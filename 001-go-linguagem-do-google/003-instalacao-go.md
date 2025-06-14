

# 03 Instalação do Go

## Resumo
Nesta aula, abordamos a instalação da linguagem de programação Go. O primeiro passo é baixar a versão mais recente no site oficial e, após a instalação, verificar se tudo está funcionando corretamente com o comando go version no terminal.Discutimos também o conceito de Go Workspace, que é onde ficará todo o código Go. Essa pasta deve ser chamada go e conter três subpastas: bin (para arquivos compilados), pkg (para pacotes compartilhados) e src (para o código fonte).Além disso, foi recomendado o uso do Visual Studio Code como editor de texto, com a instalação da extensão Go para facilitar o desenvolvimento. Por fim, foi mencionado que no próximo vídeo iniciaremos a criação do nosso primeiro programa em Go, o Hello World.

## Transcrição

Se queremos aprender a trabalhar com a linguagem de programação Go, primeiramente devemos instalá-la. Para isso, no site da linguagem (https://golang.org/), há um link de download (https://golang.org/dl/).

Então, basta fazer o download e instalar a última versão da linguagem para o nosso sistema operacional.

    Para saber mais, disponibilizamos este exercício com mais detalhes da instalação para cada sistema operacional.

Ao término da instalação, o Go será habilitado para ser utilizado na linha de comando. Para verificar a sua instalação, executamos no terminal/linha de comando:

go version

Executando o comando acima, teremos uma saída parecida com esta:

go version go1.8.1 linux/amd64

Mas além da instalação do Go, devemos seguir outros padrões para trabalhar com essa linguagem.
Go Workspace

No momento da instalação é definido um local onde ficará todo o nosso código Go, esse local é o Go Workspace. Por padrão essa pasta deve ser chamada go e deve ficar na raiz do nosso usuário. Isso vale para todos os sistemas operacionais, ou seja, na pasta do nosso usuário, nós criamos a pasta go.

Além disso, o Go Workspace deve possuir três pastas. A primeira delas é a pkg, onde ficará os pacotes compartilhados das nossas aplicações, pois o Go é uma linguagem bastante modular, dependendo bastante de pacotes que vamos importando no nosso código, veremos isso mais à frente. Além da pasta pkg, também deve ter a pasta src, onde escreveremos o código fonte de cada aplicação, e a pasta bin, onde ficará os compilados do nosso código Go.

A estrutura final do Go Workspace ficará assim:

pasta-do-usuario/
└── go
    ├── bin
    ├── pkg
    └── src

Mas se quisermos que o Go Workspace seja em outro lugar, pode? Pode ser, mas por padrão o instalador do Go já espera que essa pasta esteja na raiz do usuário, então, para nos facilitar, vamos seguir os padrões da linguagem.
Editor de texto

Mas além disso, precisamos de um editor de texto que nos ajude a codificar em Go. Um editor de texto que é bastante integrado com a linguagem, que será utilizado no treinamento, é o Visual Studio Code.

Sabendo disso, podemos fazer o nosso primeiro programa em Go, o famoso Hello World. Então, dentro da pasta go/src, vamos criar a pasta hello. E através do Visual Studio Code, nós abrimos essa pasta e dentro dela criamos o programa hello.go.

Assim que criamos o arquivo, recebemos uma notificação do Visual Studio Code: The 'Go' extension is recommended for this file type. Ou seja, existem extensões do editor que nos facilita o trabalho com a linguagem. Ao clicar em Show Recommendations, nós vemos essas extensões, e uma delas é a Go, do autor lukehoban.

Nós podemos instalar essa extensão, clicando em Install e logo em seguida em Reload, para recarregarmos a janela. No momento em que a janela for recarregada, a extensão será habilitada, mas ela tem algumas dependências, conforme é exibida na notificação do Visual Studio Code. Então, basta clicar em Install All, para instalar todas as dependências, que facilitarão a nossa codificação em Go.

Por isso é uma recomendação utilizar o Visual Studio para programar Go, pois ele facilita muito o nosso processo. Algumas funcionalidade desse editor serão exploradas ao londo do treinamento

Após o término da instalação das dependências, no próximo vídeo vamos dar início ao nosso primeiro programa em Go, o Hello World.





# Download and install

<https://go.dev/doc/install>
Download and install Go quickly with the steps described here.

For other content on installing, you might be interested in:

    Managing Go installations -- How to install multiple versions and uninstall.
    Installing Go from source -- How to check out the sources, build them on your own machine, and run them.

Download (1.24.4)
Go installation¶

Select the tab for your computer's operating system below, then follow its installation instructions.

    Remove any previous Go installation by deleting the /usr/local/go folder (if it exists), then extract the archive you just downloaded into /usr/local, creating a fresh Go tree in /usr/local/go:

~~~~bash
          $ rm -rf /usr/local/go && tar -C /usr/local -xzf go1.24.4.linux-amd64.tar.gz
          
~~~~

(You may need to run each command separately with the necessary permissions, as root or through sudo.)

Do not untar the archive into an existing /usr/local/go tree. This is known to produce broken Go installations.
Add /usr/local/go/bin to the PATH environment variable.

You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):
```bash
          export PATH=$PATH:/usr/local/go/bin
          
```

Note: Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as source $HOME/.profile.
Verify that you've installed Go by opening a command prompt and typing the following command:
```bash
          $ go version
          
```

    Confirm that the command prints the installed version of Go.

You're all set!¶

Visit the Getting Started tutorial to write some simple Go code. It takes about 10 minutes to complete. 






### Repositório Git

<https://github.com/fernandomullerjr/formacao-go-alura>

Perfeito! Se você já tem uma **pasta local com arquivos existentes** e quer transformar essa pasta em um repositório Git, vinculando ao repositório remoto `git@github.com:fernandomullerjr/formacao-go-alura.git`, aqui está o passo a passo detalhado dos comandos:

---

## ✅ **Comandos para inicializar e conectar um repositório Git em uma pasta existente**

```bash
# 1. Acesse a pasta que você quer versionar
cd /caminho/da/sua/pasta

# 2. Inicialize o repositório Git na pasta
git init

# 3. Adicione o repositório remoto (se ainda não estiver configurado)
git remote add origin git@github.com:fernandomullerjr/formacao-go-alura.git

# 4. Adicione todos os arquivos da pasta para o controle de versão
git add .

# 5. Faça o primeiro commit
git commit -m "Primeiro commit - adicionando projeto da formação Go Alura"

# 6. (Opcional) Verifique o branch atual
git branch -M main  # Se quiser garantir que esteja no branch 'main'

# 7. Envie o commit para o repositório remoto
git push -u origin main
```

---

## 🔍 **Verificar se deu tudo certo**

* Veja o status:

```bash
git status
```

* Confirme se o remoto está corretamente configurado:

```bash
git remote -v
```

Deve aparecer:

```plaintext
origin	git@github.com:fernandomullerjr/formacao-go-alura.git (fetch)
origin	git@github.com:fernandomullerjr/formacao-go-alura.git (push)
```

---

## 🚩 **Possíveis problemas comuns**

* **Branch padrão:** Se o GitHub tiver o branch como `main`, garanta isso com:

```bash
git branch -M main
```

* **Chave SSH:** Verifique se sua chave SSH está adicionada ao GitHub. Caso contrário, pode aparecer erro de permissão.

---

Se quiser, posso te gerar um script `.sh` para automatizar isso também. Quer?




- Erro

~~~~bash

> git push -u origin main
To github.com:fernandomullerjr/formacao-go-alura.git
 ! [rejected]        main -> main (non-fast-forward)
error: failed to push some refs to 'github.com:fernandomullerjr/formacao-go-alura.git'
hint: Updates were rejected because the tip of your current branch is behind
hint: its remote counterpart. Integrate the remote changes (e.g.
hint: 'git pull ...') before pushing again.
hint: See the 'Note about fast-forwards' in 'git push --help' for details.

 ~/cursos/g/formacao-go-alura  main                
~~~~


Ótima pergunta, Fernando! Esse erro ocorre porque o repositório remoto já possui algum commit (provavelmente aquele README inicial, `.gitignore`, ou licença que você criou pelo GitHub).

## 🚩 **Solução Segura – Passo a Passo**

### ✅ Opção 1 — Mesclar o que está remoto com seu local (🟢 Recomendada se não quiser perder nada)

1. **Puxe o que está no repositório remoto:**

```bash
git pull origin main --allow-unrelated-histories
```

> 🔸 O parâmetro `--allow-unrelated-histories` é necessário quando são dois históricos Git diferentes (o da sua pasta local e o do GitHub).

2. Resolva qualquer conflito, se houver (geralmente não há se foi só um README).

3. Depois envie normalmente:

```bash
git push -u origin main
```

---

### ✅ Opção 2 — Sobrescrever o remoto com o que está local (⚠️ Perigoso, sobrescreve tudo no GitHub)

Se você quer forçar que o conteúdo da sua pasta local **substitua tudo no GitHub**, faça:

```bash
git push -u origin main --force
```

> ⚠️ Isso sobrescreve todo o histórico remoto. Use com cuidado.

---

## 🔍 **Checklist antes de executar**

* Você quer preservar o conteúdo remoto? → Use **Opção 1 (pull)**.
* Você quer sobrescrever o remoto com o que está na sua pasta local? → Use **Opção 2 (force push)**.

---
 





## Instalação
<https://go.dev/dl/go1.24.4.linux-amd64.tar.gz>

wget https://go.dev/dl/go1.24.4.linux-amd64.tar.gz


## Troubleshooting 

- Problemas

> wget https://go.dev/dl/go1.24.4.linux-amd64.tar.gz
--2025-06-14 14:41:50--  https://go.dev/dl/go1.24.4.linux-amd64.tar.gz
Resolving go.dev (go.dev)... 2001:4860:4802:32::15, 2001:4860:4802:34::15, 2001:4860:4802:36::15, ...
Connecting to go.dev (go.dev)|2001:4860:4802:32::15|:443.


- tshoot

Perfeito. O seu erro tem um detalhe importante: o WSL está tentando resolver endereços **IPv6** (`2001:4860...`) e provavelmente **seu ambiente não tem suporte a IPv6 habilitado ou funcional**.

Isso é **extremamente comum no WSL2**, pois o stack de rede do Windows muitas vezes tem suporte parcial ou quebrado para IPv6 dentro do WSL.

---

# 🩹 **Solução prática e imediata — Desabilitar IPv6 no WSL**

### ✅ Crie o arquivo (ou edite) `/etc/wsl.conf`:

```bash
sudo nano /etc/wsl.conf
```

Adicione:

```ini
[network]
generateResolvConf = false
```

Salve.

---

### ✅ Edite o resolv.conf manualmente:

```bash
sudo rm -f /etc/resolv.conf
sudo bash -c 'echo "nameserver 8.8.8.8" > /etc/resolv.conf'
sudo bash -c 'echo "nameserver 8.8.4.4" >> /etc/resolv.conf'
```

**Proteja o arquivo:**

```bash
sudo chattr +i /etc/resolv.conf
```

---

### ✅ Desabilite IPv6 no WSL temporariamente (teste rápido):

Execute dentro do WSL:

```bash
sudo sysctl -w net.ipv6.conf.all.disable_ipv6=1
sudo sysctl -w net.ipv6.conf.default.disable_ipv6=1
```

Teste:

```bash
ping google.com
```

ou

```bash
wget https://go.dev/dl/go1.22.4.linux-amd64.tar.gz
```

Se funcionar, você pode tornar isso permanente.

---

### ✅ Para desabilitar IPv6 permanentemente:

1. Edite:

```bash
sudo nano /etc/sysctl.conf
```

Adicione no final:

```plaintext
net.ipv6.conf.all.disable_ipv6 = 1
net.ipv6.conf.default.disable_ipv6 = 1
net.ipv6.conf.lo.disable_ipv6 = 1
```

Salve.

2. Aplique:

```bash
sudo sysctl -p
```

---

## 🔥 Reinicie o WSL:

No **PowerShell (Admin)**:

```powershell
wsl --shutdown
```

Abra novamente seu Ubuntu.

---

## 🚀 Teste novamente:

```bash
ping google.com
wget https://go.dev/dl/go1.22.4.linux-amd64.tar.gz
```

---

# 🏆 **Se ainda assim não funcionar:**

Me fala que eu te ensino como fazer um download manual via navegador do Windows e depois mover o arquivo para dentro do WSL usando:

```bash
mv /mnt/c/Users/SEU_USUARIO/Downloads/go1.22.4.linux-amd64.tar.gz .
```

E daí seguimos normalmente a instalação.

---

Quer que eu te mande esse procedimento já convertido num script `.sh` para executar automático? 💪




- Desativando ipv6, funcionou o download


>
sudo sysctl -w net.ipv6.conf.all.disable_ipv6=1
sudo sysctl -w net.ipv6.conf.default.disable_ipv6=1
net.ipv6.conf.all.disable_ipv6 = 1
net.ipv6.conf.default.disable_ipv6 = 1
>
>
>
>
> wget https://go.dev/dl/go1.24.4.linux-amd64.tar.gz
--2025-06-14 14:45:28--  https://go.dev/dl/go1.24.4.linux-amd64.tar.gz
Resolving go.dev (go.dev)... 216.239.36.21, 216.239.38.21, 216.239.32.21, ...
Connecting to go.dev (go.dev)|216.239.36.21|:443... connected.
HTTP request sent, awaiting response... 302 Found
Location: https://dl.google.com/go/go1.24.4.linux-amd64.tar.gz [following]
--2025-06-14 14:45:29--  https://dl.google.com/go/go1.24.4.linux-amd64.tar.gz
Resolving dl.google.com (dl.google.com)... 172.217.30.46, 2800:3f0:4001:83b::200e
Connecting to dl.google.com (dl.google.com)|172.217.30.46|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 78559214 (75M) [application/x-gzip]
go1.24.4.linux-amd64.tar.gz: Permission denied

Cannot write to ‘go1.24.4.linux-amd64.tar.gz’ (Success).
> ls
cni  config  containerd  k8sgpt_amd64.deb
> sudo wget https://go.dev/dl/go1.24.4.linux-amd64.tar.gz
--2025-06-14 14:45:39--  https://go.dev/dl/go1.24.4.linux-amd64.tar.gz
Resolving go.dev (go.dev)... 216.239.32.21, 216.239.38.21, 216.239.36.21, ...
Connecting to go.dev (go.dev)|216.239.32.21|:443... connected.
HTTP request sent, awaiting response... 302 Found
Location: https://dl.google.com/go/go1.24.4.linux-amd64.tar.gz [following]
--2025-06-14 14:45:39--  https://dl.google.com/go/go1.24.4.linux-amd64.tar.gz
Resolving dl.google.com (dl.google.com)... 172.217.30.46, 2800:3f0:4001:83b::200e
Connecting to dl.google.com (dl.google.com)|172.217.30.46|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 78559214 (75M) [application/x-gzip]
Saving to: ‘go1.24.4.linux-amd64.tar.gz’

go1.24.4.linux-amd64.tar.gz   100%[=================================================>]  74.92M   483KB/s    in 1m 54s

2025-06-14 14:47:28 (674 KB/s) - ‘go1.24.4.linux-amd64.tar.gz’ saved [78559214/78559214]

> ls
cni  config  containerd  go1.24.4.linux-amd64.tar.gz  k8sgpt_amd64.deb
> date
Sat Jun 14 14:49:51 -03 2025

 !w /opt       



 
 
 
tem algum jeito de deixar funcionando o ipv6 com wsl2 e ubuntu 22?



Perfeito, Fernando. Esse é um ponto muito relevante, e a sua pergunta é extremamente válida: **sim, dá para ativar e fazer IPv6 funcionar no WSL2, mas não é suportado oficialmente de forma completa ainda.** O suporte vem sendo desenvolvido nas versões mais recentes do Windows (Windows 11 especialmente), mas ainda depende de configuração manual.

---

# 🌐 **✔️ Como habilitar IPv6 no WSL2 no Ubuntu 22 — Guia Definitivo**

## 🚨 Atenção:

* O WSL2 faz NAT sobre o Windows Host, que por padrão **não repassa IPv6.**
* Você pode habilitar IPv6 de forma **local (dentro da máquina)** ou, com ajustes mais complexos, até roteá-lo corretamente para fora (internet IPv6 pública), se sua rede/operadora permitir.

---

## ✅ **Passo 1 — Verificar se o Windows tem IPv6 ativo**

1. No Windows, abra o **Prompt de comando** ou **PowerShell**:

```powershell
ping -6 google.com
```

✔️ Se responder, significa que o Windows tem IPv6 funcionando.

Se **não responder**, vá para:

* **Painel de Controle → Rede e Internet → Alterar Configurações do Adaptador**
* Clique com o botão direito na sua placa de rede → **Propriedades**
* ✔️ Marque **"Protocolo da Internet versão 6 (TCP/IPv6)"**

---

## ✅ **Passo 2 — Verificar se WSL2 está com suporte interno a IPv6**

1. Dentro do Ubuntu WSL:

```bash
ip a
```

Veja se aparece uma linha do tipo:

```plaintext
inet6 fe80::.../64 scope link
```

Se tem apenas `fe80::` (link-local), significa que **não há roteamento IPv6 para fora, apenas local**.

---

## ✅ **Passo 3 — Habilitar IPv6 manualmente dentro do WSL**

1. Crie ou edite `/etc/sysctl.conf`:

```bash
sudo nano /etc/sysctl.conf
```

Adicione (ou descomente) as linhas:

```plaintext
net.ipv6.conf.all.disable_ipv6 = 0
net.ipv6.conf.default.disable_ipv6 = 0
net.ipv6.conf.lo.disable_ipv6 = 0
```

Salve.

2. Aplique:

```bash
sudo sysctl -p
```

3. Teste:

```bash
ping6 google.com
```

### 👉 Se **não funcionar**, continue.

---

## ✅ **Passo 4 — Habilitar Encaminamento IPv6 no Windows para o WSL**

1. No **PowerShell (Admin)**, execute:

```powershell
Get-NetAdapter | Where-Object {$_.Name -like "*WSL*"}
```

Pegue o nome do adaptador WSL (`vEthernet (WSL)` geralmente).

2. Habilite o IPv6 Router Advertisement:

```powershell
Set-NetAdapter -Name "vEthernet (WSL)" -RouterDiscovery Enabled
```

3. Ative o Encaminamento de Pacotes IPv6 no registro do Windows:

```powershell
New-ItemProperty -Path "HKLM:\SYSTEM\CurrentControlSet\Services\Tcpip6\Parameters" -Name "DisabledComponents" -PropertyType DWord -Value 0
```

4. Ou ative via regedit:

* Vá em `HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\Tcpip6\Parameters`
* Crie (ou edite) um valor DWORD `DisabledComponents` e coloque `0`.

---

## ✅ **Passo 5 — Fazer NAT ou bridge manual para IPv6**

> ⚙️ Este é o passo avançado. O WSL2 não faz bridge nativa, então você pode criar um bridge manual usando **ICS (Internet Connection Sharing)** ou um **proxy NAT IPv6**.

### Ferramenta mais comum para isso:

* [**Wsl2-network-fix**](https://github.com/sakai135/wsl2-network-fix) → Script do GitHub que faz bridge para WSL2 (incluindo IPv6 experimental).

---

## ✅ **Passo 6 — Alternativa simples (meio gambiarra, mas funciona para resolver o DNS IPv6 quebrado):**

1. Forçar o DNS a resolver via IPv4 e não tentar IPv6:

```bash
sudo nano /etc/gai.conf
```

Descomente essa linha:

```plaintext
precedence ::ffff:0:0/96  100
```

✔️ Isso faz com que o WSL **priorize IPv4 sobre IPv6**, mesmo que o IPv6 esteja tecnicamente ativado.

---

## 🚀 **Resumo prático:**

* ✔️ **IPv6 local:** WSL responde endereços `fe80::` (link-local), útil para comunicação entre containers e com o host.
* ✔️ **IPv6 externo:** Só funciona com roteamento avançado, NAT66 ou bridge manual via ferramentas como `wsl2-network-fix`.
* ✔️ **Caminho oficial:** A Microsoft está desenvolvendo suporte nativo melhor para IPv6 no WSL, já disponível em insiders e em testes no Windows 11 24H2.

---

## 🏆 **Se quiser, te entrego:**

* 🔧 Um script PowerShell que habilita tudo no Windows.
* 🐧 Um script bash que configura o lado Linux (Ubuntu WSL).
* 📜 Documentação customizada passo a passo.

Me responde: **quer que eu gere esses scripts e te entregue pronto?** 💪






- Ir na linha de desativar ipv6 permanentemente


### ✅ Para desabilitar IPv6 permanentemente:

1. Edite:

```bash
sudo vi /etc/sysctl.conf
```

Adicione no final:

```plaintext
net.ipv6.conf.all.disable_ipv6 = 1
net.ipv6.conf.default.disable_ipv6 = 1
net.ipv6.conf.lo.disable_ipv6 = 1
```

Salve.

2. Aplique:

```bash
sudo sysctl -p
```

---

## 🔥 Reinicie o WSL:

No **PowerShell (Admin)**:

```powershell
wsl --shutdown
```

Abra novamente seu Ubuntu.

---

## 🚀 Teste novamente:

```bash
ping google.com
wget https://go.dev/dl/go1.22.4.linux-amd64.tar.gz
```




## PENDENTE

- RETOMAR INSTALAÇÃO DO GO
- Depois que a rede funcionar, você pode seguir a instalação do Go normalmente:

~~~~bash
cd /opt
sudo wget https://go.dev/dl/go1.22.4.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
go version
~~~~





- RETOMAR INSTALAÇÃO DO GO

> ls
cni  config  containerd  go1.24.4.linux-amd64.tar.gz  k8sgpt_amd64.deb

~~~~bash
cd /opt
sudo tar -C /usr/local -xzf go1.24.4.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
go version
~~~~


Perfeita sua dúvida, Fernando. A resposta é: **sim, funciona — mas existe uma maneira mais adequada quando você usa `zsh` (ou qualquer shell específico como bash, fish, etc.).**

---

# 🛠️ **Explicação Técnica — Válido para `zshrc`?**

## ✅ Esse trecho do material está correto funcionalmente, **mas aplica a configuração no nível global do sistema**, porque altera o `/etc/profile`.

✔️ Isso afeta qualquer shell que use esse arquivo, incluindo sessões bash, zsh e outros, **mas só nas sessões de login interativo ou shells que leem o `/etc/profile`.**

⚠️ No entanto, o `zsh` **não lê diretamente o `/etc/profile`** no dia a dia. Ele lê:

* `~/.zshrc` → quando abre o terminal interativo.
* `~/.zprofile` → em sessões de login.

Portanto, se você quer que a configuração do Go seja aplicada **apenas para seu usuário e no zsh**, o melhor caminho é:

---

# ✅ **💯 Forma recomendada para quem usa ZSH**

### ✔️ Adicione no arquivo `~/.zshrc`:

```bash
export PATH=$PATH:/usr/local/go/bin
```

> 🔥 Isso faz com que, sempre que você abrir um terminal zsh, o Go esteja disponível.

### ⚙️ Aplique imediatamente:

```bash
source ~/.zshrc
```

### ✔️ Teste:

```bash
go version
```

---

## ⚠️ **Comparação entre os métodos:**

| Método                        | Onde configura   | Quem afeta                          | Quando usar                                   |
| ----------------------------- | ---------------- | ----------------------------------- | --------------------------------------------- |
| `/etc/profile`                | Sistema todo     | Todos os usuários e shells de login | ✅ Útil em ambientes multi-usuário ou servidor |
| `~/.bashrc` ou `~/.zshrc`     | Usuário atual    | Apenas aquele shell (bash ou zsh)   | ✅ Mais recomendado para uso pessoal           |
| `~/.profile` ou `~/.zprofile` | Sessões de login | Bash e Zsh em login                 | Menos comum, mas válido para certos ambientes |

---

## 🏆 **✔️ Meu conselho para quem usa WSL + Ubuntu + ZSH:**

✅ Melhor prática:

```bash
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
source ~/.zshrc
```

Se você quiser também garantir para todos os shells (incluindo bash, zsh, scripts que rodam fora do terminal), pode opcionalmente adicionar em `/etc/profile`:

```bash
sudo nano /etc/profile
```

E adicionar:

```bash
export PATH=$PATH:/usr/local/go/bin
```

---



Perfeito, Fernando! Aqui está um script completo em **Bash**, que você pode rodar diretamente no seu terminal do **WSL + Ubuntu + ZSH** para:

1. Fazer o download da versão mais recente do Go (ou a que você desejar).
2. Instalar no diretório `/usr/local`.
3. Configurar automaticamente o PATH no seu arquivo `~/.zshrc`.
4. Verificar se o Go está instalado corretamente.

---

# 🚀 **Script completo – Instalação do Go + Configuração do ZSH**

### 🔧 Salve este conteúdo como, por exemplo, `instalar-go.sh` e execute com:

```bash
chmod +x instalar-go.sh
./instalar-go.sh
```

---

### 🧠 **Conteúdo do script:**

```bash
#!/bin/bash

# ===============================
# Script de instalação do Go no Ubuntu (WSL) com configuração para ZSH
# Autor: Fernando Muller + ChatGPT
# ===============================

# ===== CONFIGURAÇÃO =====
GO_VERSION="1.22.4"
GO_TAR_FILE="go${GO_VERSION}.linux-amd64.tar.gz"
GO_DOWNLOAD_URL="https://go.dev/dl/${GO_TAR_FILE}"
INSTALL_DIR="/usr/local"
ZSHRC_FILE="$HOME/.zshrc"

# =========================

echo "🚀 Iniciando instalação do Go versão ${GO_VERSION}..."

# ✅ Removendo instalações anteriores
sudo rm -rf ${INSTALL_DIR}/go

# ✅ Baixando o Go
echo "⬇️  Baixando ${GO_TAR_FILE}..."
wget -c ${GO_DOWNLOAD_URL} -O ${GO_TAR_FILE}

if [ $? -ne 0 ]; then
    echo "❌ Erro no download do Go. Verifique sua conexão ou o link."
    exit 1
fi

# ✅ Instalando
echo "📦 Extraindo para ${INSTALL_DIR}..."
sudo tar -C ${INSTALL_DIR} -xzf ${GO_TAR_FILE}

# ✅ Configurando PATH no .zshrc
if grep -q "/usr/local/go/bin" ${ZSHRC_FILE}; then
    echo "✅ PATH do Go já configurado no ${ZSHRC_FILE}"
else
    echo "🔧 Adicionando Go ao PATH no ${ZSHRC_FILE}..."
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ${ZSHRC_FILE}
fi

# ✅ Aplicando as mudanças no shell atual
source ${ZSHRC_FILE}

# ✅ Verificando instalação
echo "🧠 Verificando instalação do Go..."
go version

if [ $? -eq 0 ]; then
    echo "🎉 Go instalado e configurado com sucesso!"
else
    echo "⚠️ Algo deu errado na instalação do Go."
fi

# ✅ Limpando o arquivo .tar.gz
rm -f ${GO_TAR_FILE}

echo "✅ Script finalizado."
```

---

# 💡 **Como executar:**

1. Crie o arquivo:

```bash
nano instalar-go.sh
```

Cole o conteúdo.

2. Dê permissão de execução:

```bash
chmod +x instalar-go.sh
```

3. Rode:

```bash
./instalar-go.sh
```

---

## ✔️ Resultado esperado no final:

```plaintext
go version go1.22.4 linux/amd64
🎉 Go instalado e configurado com sucesso!
```

---





> /home/fernando/cursos/golang/formacao-go-alura/001-go-linguagem-do-google/003-instalador-go-automatizado.sh
🚀 Iniciando instalação do Go versão 1.24.4...
⬇️  Baixando go1.24.4.linux-amd64.tar.gz...
go1.24.4.linux-amd64.tar.gz: Permission denied
❌ Erro no download do Go. Verifique sua conexão ou o link.
> sudo /home/fernando/cursos/golang/formacao-go-alura/001-go-linguagem-do-google/003-instalador-go-automatizado.sh
🚀 Iniciando instalação do Go versão 1.24.4...
⬇️  Baixando go1.24.4.linux-amd64.tar.gz...
--2025-06-14 14:59:05--  https://go.dev/dl/go1.24.4.linux-amd64.tar.gz
Resolving go.dev (go.dev)... 216.239.34.21, 216.239.38.21, 216.239.32.21, ...
Connecting to go.dev (go.dev)|216.239.34.21|:443... connected.
HTTP request sent, awaiting response... 302 Found
Location: https://dl.google.com/go/go1.24.4.linux-amd64.tar.gz [following]
--2025-06-14 14:59:06--  https://dl.google.com/go/go1.24.4.linux-amd64.tar.gz
Resolving dl.google.com (dl.google.com)... 172.217.30.46, 2800:3f0:4001:83b::200e
Connecting to dl.google.com (dl.google.com)|172.217.30.46|:443... connected.
HTTP request sent, awaiting response... 416 Requested range not satisfiable

    The file is already fully retrieved; nothing to do.

📦 Extraindo para /usr/local...
🔧 Adicionando Go ao PATH no /root/.zshrc...
🧠 Verificando instalação do Go...
go version go1.24.4 linux/amd64
🎉 Go instalado e configurado com sucesso!
✅ Script finalizado.
> date
Sat Jun 14 14:59:14 -03 2025

 !w /opt    






> go version
zsh: command not found: go

 !w /opt                



 Perfeito, Fernando! Esse erro (`zsh: command not found: go`) indica que o Go foi instalado, mas o **`PATH` não está corretamente configurado** no seu ambiente atual.

Vamos validar e resolver.

---

## ✅ **Passo 1 — Verificar se o Go está instalado no lugar certo**

Execute:

```bash
ls /usr/local/go
```

Se retornar uma lista como:

```plaintext
bin  doc  lib  misc  pkg  src
```

✔️ O Go está instalado corretamente na pasta `/usr/local/go`.

---

## ✅ **Passo 2 — Verificar se o PATH está correto no ZSH**

Execute:

```bash
echo $PATH
```

Verifique se aparece `/usr/local/go/bin`.

Se **não aparece**, é porque a linha não foi corretamente adicionada no `~/.zshrc`.

---

## ✅ **Passo 3 — Adicionar o Go no PATH manualmente**

1. Abra seu arquivo de configuração do ZSH:

```bash
vi ~/.zshrc
```

2. Adicione no final do arquivo:

```bash
export PATH=$PATH:/usr/local/go/bin
```

3. Salve (`Ctrl + O` → Enter → `Ctrl + X`).

4. Aplique a mudança:

```bash
source ~/.zshrc
```

---

## ✅ **Passo 4 — Testar**

```bash
go version
```

✔️ Resultado esperado:

```plaintext
go version go1.22.4 linux/amd64
```

---

## 🚀 **Validação adicional (Se quiser testar na hora, sem editar arquivo)**

Execute no terminal diretamente:

```bash
export PATH=$PATH:/usr/local/go/bin
go version
```

Se funcionar, é confirmação total de que o problema era apenas o PATH não carregado.

---

## 🔥 **DICA FINAL – Para WSL**

Se você quiser que isso funcione também para qualquer shell (bash, zsh, etc.) e qualquer terminal no WSL, pode adicionar também em:

```bash
sudo vi /etc/profile
```

Adicione no final:

```bash
export PATH=$PATH:/usr/local/go/bin
```

Salve e rode:

```bash
source /etc/profile
```

Assim, estará disponível em todos os shells do WSL para todos os usuários.

---

Se quiser, posso te gerar um script que:

* Valida se já tem Go.
* Adiciona no `.zshrc` automaticamente se não tiver.
* Faz o reload automático do terminal.

Me confirma se quer! 💪





> go version
go version go1.24.4 linux/amd64
> date
Sat Jun 14 15:04:41 -03 2025
