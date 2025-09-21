12
Mãos na Massa: Delay e Aumentando o número de testes
Próxima Atividade

Agora vamos fazer com que a cada vez que o usuário selecione a opção de monitoramento, cada site seja testado mais de uma vez, de acordo com o que o usuário setar nas constantes.

1- Primeiro crie a constante monitoramentos que indicará quantas vezes o site será testado. Coloque o número que desejar, mas lembre-se que quanto mais você testar, mais o script vai demorar:

//hello.go
package main

import (
    "fmt"
    "net/http"
    "os"    
)

const monitoramentos = 3

// restante do arquivo

2- Agora vamos colocar um for dentro da função iniciarMonitoramentos para que o slice seja percorrido o número de vezes que configuramos na constante:

//hello.go

//restante do arquivo

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://httpbin.org/status/200",
        "https://www.alura.com.br", "https://www.caelum.com.br"}

    //Adição aqui
    for i := 0; i < monitoramentos; i++ {
        for i, site := range sites {
            fmt.Println("Testando site", i, ":", site)
            testaSite(site)
        }
    }

    fmt.Println("")
}

3- Para que haja uma pausa entre cada um dos testes que faremos, vamos adicionar um pequeno delay entre cada iteração de monitoramento. Utilize a função Sleep do pacote time para dar uma pausa entre monitoramentos. Para que o tamanho desta pausa fique configurável, vamos criar uma constante que vai dizer o tamanho do delay:

//hello.go
package main

import (
    "fmt"
    "net/http"
    "os"    
)

const monitoramentos = 3
const delay = 5

// restante do arquivo

4- E a partir do valor da constante, vamos adicionar um time.Sleep na função iniciarMonitoramento. Vamos multiplicar a nossa contante (delay) pela constante que representa o número de segundos no Go(time.Seconds):

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://httpbin.org/status/200",
        "https://www.alura.com.br", "https://www.caelum.com.br"}

    for i := 0; i < monitoramentos; i++ {
        for i, site := range sites {
            fmt.Println("Testando site", i, ":", site)
            testaSite(site)
        }

        // adição AQUI!
        time.Sleep(delay * time.Second)
        fmt.Println("")
    }
    fmt.Println("")
}

_ Adicione o outro fmt.Println("") para que a exibição fique mais organizada para o seu usuário_

Agora o seu script deve estar testando os seus sites de acordo com o número de vezes que você setou na constante monitoramentos, e entre cada testada ele deve dar uma pausa de acordo com o número de segundos que você configurou na constante delay.







# 12 - Mãos na Massa: Delay e Aumentando o número de testes

## 🎯 **Objetivo**
Fazer com que a cada vez que o usuário selecione a opção de monitoramento, cada site seja testado múltiplas vezes, de acordo com as constantes configuradas.

---

## 📋 **Passo a Passo**

### **1. Criando a Constante de Monitoramentos**
Primeiro crie a constante `monitoramentos` que indicará quantas vezes o site será testado:

```go
// hello.go
package main

import (
    "fmt"
    "net/http"
    "os"    
)

const monitoramentos = 3

// restante do arquivo
```

> ⚠️ **Atenção**: Quanto mais você testar, mais o script vai demorar!

---

### **2. Implementando Loop Aninhado**
Coloque um `for` dentro da função `iniciarMonitoramento` para que o slice seja percorrido o número de vezes configurado:

```go
// hello.go
// restante do arquivo

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{
        "https://httpbin.org/status/200",
        "https://www.alura.com.br", 
        "https://www.caelum.com.br"
    }

    // Adição aqui
    for i := 0; i < monitoramentos; i++ {
        for i, site := range sites {
            fmt.Println("Testando site", i, ":", site)
            testaSite(site)
        }
    }

    fmt.Println("")
}
```

---

### **3. Criando Constante de Delay**
Para adicionar uma pausa entre cada teste, crie uma constante configurável:

```go
// hello.go
package main

import (
    "fmt"
    "net/http"
    "os"
    "time"    
)

const monitoramentos = 3
const delay = 5

// restante do arquivo
```

---

### **4. Implementando o Delay**
Adicione `time.Sleep` na função `iniciarMonitoramento`, multiplicando a constante `delay` por `time.Second`:

```go
func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{
        "https://httpbin.org/status/200",
        "https://www.alura.com.br", 
        "https://www.caelum.com.br"
    }

    for i := 0; i < monitoramentos; i++ {
        for i, site := range sites {
            fmt.Println("Testando site", i, ":", site)
            testaSite(site)
        }

        // Adição AQUI!
        time.Sleep(delay * time.Second)
        fmt.Println("")
    }
    fmt.Println("")
}
```

> 💡 **Dica**: Adicione o `fmt.Println("")` extra para que a exibição fique mais organizada para o usuário.

---

## ✅ **Resultado**
Agora o script deve estar:
- ✅ Testando os sites de acordo com o número configurado em `monitoramentos`
- ✅ Fazendo pausa entre cada rodada conforme o `delay` em segundos
- ✅ Exibindo informações organizadas para o usuário

Similar code found with 1 license type