


https://go.dev/doc/install


<https://go.dev/doc/install>


Download and install

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
