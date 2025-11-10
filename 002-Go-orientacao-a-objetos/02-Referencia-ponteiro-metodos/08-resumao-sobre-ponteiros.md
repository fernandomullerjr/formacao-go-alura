# Resumão: Referência, Ponteiros e Métodos em Go

## 📚 Conceitos Fundamentais

### 1. Structs em Go
```go
type ContaCorrente struct {
    titular       string
    numeroAgencia int
    numeroConta   int
    saldo         float64
}
```

### 2. Formas de Criar Structs

#### Sintaxe com Campos Nomeados
```go
conta := ContaCorrente{
    titular: "João",
    numeroAgencia: 123,
    numeroConta: 456789,
    saldo: 100.0
}
```

#### Sintaxe Curta (Ordem dos Campos)
```go
conta := ContaCorrente{"João", 123, 456789, 100.0}
```

## 🎯 Ponteiros: Conceitos Essenciais

### O que são Ponteiros?
- **Ponteiro**: Uma variável que armazena o **endereço de memória** de outra variável
- **Analogia**: Se uma struct é um "apartamento", o ponteiro é o "endereço do apartamento"

### Operadores de Ponteiro
- **`&`**: Operador de referência (obtém o endereço)
- **`*`**: Operador de desreferência (obtém o valor no endereço)

```go
var x int = 42
var ptr *int = &x  // ptr recebe o endereço de x
fmt.Println(ptr)   // endereço: 0xc000014098
fmt.Println(*ptr)  // valor: 42
```

## 🏗️ Criando Structs com Ponteiros

### Método 1: new()
```go
var conta *ContaCorrente
conta = new(ContaCorrente)
conta.titular = "Maria"
conta.saldo = 500.0
```

### Método 2: &struct{}
```go
conta := &ContaCorrente{
    titular: "Pedro",
    saldo: 300.0,
}
```

## ⚖️ Comparações: Valores vs Ponteiros

### Comparando Structs (Valores)
```go
conta1 := ContaCorrente{"Ana", 123, 456, 100}
conta2 := ContaCorrente{"Ana", 123, 456, 100}
fmt.Println(conta1 == conta2) // true - compara CONTEÚDO
```

### Comparando Ponteiros
```go
var ptr1 *ContaCorrente = new(ContaCorrente)
ptr1.titular = "Carlos"

var ptr2 *ContaCorrente = new(ContaCorrente) 
ptr2.titular = "Carlos"

fmt.Println(ptr1 == ptr2)   // false - compara ENDEREÇOS
fmt.Println(*ptr1 == *ptr2) // true - compara CONTEÚDO
```

## 🔧 Métodos em Structs

### Método com Receptor por Valor
```go
func (c ContaCorrente) ExibirSaldo() {
    fmt.Printf("Saldo: R$ %.2f\n", c.saldo)
}

// Uso:
conta := ContaCorrente{saldo: 100}
conta.ExibirSaldo() // Não modifica a struct original
```

### Método com Receptor por Ponteiro
```go
func (c *ContaCorrente) Sacar(valor float64) {
    if c.saldo >= valor {
        c.saldo -= valor
    }
}

// Uso:
conta := &ContaCorrente{saldo: 100}
conta.Sacar(50) // Modifica a struct original
```

## 💡 Quando Usar Cada Abordagem

### Use Valores Quando:
- ✅ A struct é pequena
- ✅ Você não precisa modificar os dados
- ✅ Quer evitar compartilhamento de estado

### Use Ponteiros Quando:
- ✅ A struct é grande (performance)
- ✅ Precisa modificar os dados originais
- ✅ Quer compartilhar dados entre funções
- ✅ Quer evitar cópias desnecessárias

## 🎭 Exemplos Práticos

### Sistema Bancário Completo
```go
type ContaCorrente struct {
    titular       string
    numeroAgencia int
    numeroConta   int
    saldo         float64
}

// Método que não modifica (receptor por valor)
func (c ContaCorrente) ObterSaldo() float64 {
    return c.saldo
}

// Método que modifica (receptor por ponteiro)
func (c *ContaCorrente) Depositar(valor float64) {
    c.saldo += valor
}

func (c *ContaCorrente) Sacar(valor float64) bool {
    if c.saldo >= valor {
        c.saldo -= valor
        return true
    }
    return false
}

// Uso:
func main() {
    conta := &ContaCorrente{
        titular: "João",
        numeroAgencia: 123,
        numeroConta: 456789,
        saldo: 1000.0,
    }
    
    conta.Depositar(200)        // saldo: 1200
    sucesso := conta.Sacar(300) // saldo: 900
    fmt.Println("Saldo atual:", conta.ObterSaldo())
}
```

## 🧠 Conceitos de Memória

### Visualizando Endereços
```go
conta1 := ContaCorrente{titular: "Ana"}
conta2 := &ContaCorrente{titular: "Bruno"}

fmt.Printf("conta1 (valor): %+v\n", conta1)           // {titular:Ana ...}
fmt.Printf("conta2 (ponteiro): %+v\n", conta2)        // &{titular:Bruno ...}
fmt.Printf("Endereço de conta1: %p\n", &conta1)       // 0xc000014080
fmt.Printf("Endereço que conta2 aponta: %p\n", conta2) // 0xc000014090
fmt.Printf("Endereço de conta2: %p\n", &conta2)       // 0xc000014088
```

## 🔄 Passagem de Parâmetros

### Por Valor (Cópia)
```go
func alterarSaldo(c ContaCorrente, novoSaldo float64) {
    c.saldo = novoSaldo // Não afeta a struct original
}
```

### Por Referência (Ponteiro)
```go
func alterarSaldo(c *ContaCorrente, novoSaldo float64) {
    c.saldo = novoSaldo // Modifica a struct original
}
```

## 🚨 Armadilhas Comuns

### 1. Nil Pointer Dereference
```go
var conta *ContaCorrente
// conta.titular = "erro" // PANIC: nil pointer dereference

// Correto:
conta = new(ContaCorrente)
conta.titular = "OK"
```

### 2. Comparação de Ponteiros vs Conteúdo
```go
ptr1 := &ContaCorrente{titular: "Ana"}
ptr2 := &ContaCorrente{titular: "Ana"}

fmt.Println(ptr1 == ptr2)   // false (endereços diferentes)
fmt.Println(*ptr1 == *ptr2) // true (conteúdo igual)
```

## 📋 Cheat Sheet Rápido

| Operação | Sintaxe | Resultado |
|----------|---------|-----------|
| Criar valor | `conta := ContaCorrente{}` | Struct por valor |
| Criar ponteiro | `conta := &ContaCorrente{}` | Ponteiro para struct |
| Obter endereço | `&variavel` | Ponteiro |
| Obter valor | `*ponteiro` | Valor |
| Método por valor | `func (c ContaCorrente)` | Não modifica original |
| Método por ponteiro | `func (c *ContaCorrente)` | Modifica original |

## 🎯 Regras de Ouro

1. **Performance**: Use ponteiros para structs grandes
2. **Modificação**: Use ponteiros quando precisar alterar dados
3. **Comparação**: Lembre-se da diferença entre endereço e conteúdo
4. **Nil Safety**: Sempre verifique ponteiros nil antes de usar
5. **Consistência**: Mantenha padrão consistente no projeto

## 🔚 Resumo Final

Ponteiros em Go são fundamentais para:
- **Eficiência de memória** (evitar cópias)
- **Modificação de dados** (alterar structs originais)
- **Compartilhamento** (múltiplas referências ao mesmo dado)

A escolha entre valor e ponteiro depende do contexto: use valores para dados imutáveis e pequenos, ponteiros para dados mutáveis e grandes.