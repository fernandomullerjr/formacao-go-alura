# Explicação do Código Go - Struct ContaCorrente

## Visão Geral

Este código implementa uma estrutura `ContaCorrente` em Go que representa uma conta bancária com funcionalidades básicas como saque, depósito e transferência.

## Estrutura ContaCorrente

```go
type ContaCorrente struct {
    Titular       clientes.Titular
    NumeroAgencia int
    NumeroConta   int
    Saldo         float64
}
```

### Composição
- **`Titular clientes.Titular`**: Utiliza composição para incluir dados do cliente
- **`NumeroAgencia int`**: Número da agência bancária
- **`NumeroConta int`**: Número da conta
- **`Saldo float64`**: Saldo atual da conta

## Métodos com Ponteiros (Asterisco *)

### O que é uma "Cópia" em Go?

Em Go, quando você passa uma struct **por valor** (sem ponteiro), o Go cria uma **cópia completa** de todos os dados da struct na memória. Isso significa:

```go
// Struct original na memória
contaOriginal := ContaCorrente{
    Titular:       clienteSilvia,
    NumeroAgencia: 123,
    NumeroConta:   456,
    Saldo:         1000.0,
}

// Quando passamos por VALOR, Go cria uma NOVA struct na memória
// com os mesmos valores, mas em um endereço diferente
func exemploSemPonteiro(conta ContaCorrente) {
    // 'conta' aqui é uma CÓPIA INDEPENDENTE de contaOriginal
    // Modificações aqui NÃO afetam contaOriginal
    conta.Saldo = 500.0  // Só muda a cópia, não o original
}
```

### Visualizando Cópia vs Ponteiro

```go
// === EXEMPLO COM CÓPIA (SEM PONTEIRO) ===
func (c ContaCorrente) SacarERRADO(valor float64) {
    fmt.Printf("Endereço da cópia: %p\n", &c)    // Endereço diferente!
    c.Saldo -= valor  // Modifica APENAS a cópia local
    // Quando a função termina, a cópia é descartada
}

// === EXEMPLO COM PONTEIRO (CORRETO) ===
func (c *ContaCorrente) SacarCORRETO(valor float64) {
    fmt.Printf("Endereço original: %p\n", c)     // Mesmo endereço!
    c.Saldo -= valor  // Modifica a struct ORIGINAL na memória
}

// Demonstração:
func main() {
    conta := ContaCorrente{Saldo: 1000}
    fmt.Printf("Endereço original: %p\n", &conta)  // Ex: 0xc000012040
    
    // Método ERRADO - com cópia
    conta.SacarERRADO(100)
    fmt.Println("Saldo após saque errado:", conta.Saldo)  // Ainda 1000!
    
    // Método CORRETO - com ponteiro  
    conta.SacarCORRETO(100)
    fmt.Println("Saldo após saque correto:", conta.Saldo)  // Agora 900!
}
```

### Por que usar ponteiros?

O **asterisco (*)** indica que estamos trabalhando com **ponteiros**. Em Go, quando você quer que um método modifique o valor original de uma struct (não uma cópia), você precisa usar ponteiros.

```go
func (c *ContaCorrente) Sacar(valorDoSaque float64) string
//    ^ Este asterisco indica que 'c' é um ponteiro para ContaCorrente
```

**Sem ponteiro (valor) - Trabalha com CÓPIA:**
- Go faz uma **cópia completa** da struct na memória
- A cópia tem um endereço de memória diferente
- Modificações afetam apenas a cópia local
- Quando a função termina, a cópia é descartada
- A struct original permanece inalterada

**Com ponteiro (referência) - Trabalha com ORIGINAL:**
- Go trabalha diretamente com a struct original na memória
- Usa o mesmo endereço de memória
- Modificações **alteram** a struct original
- Necessário para operações que modificam estado

### Exemplo Prático Detalhado

```go
type ContaCorrente struct {
    Saldo float64
}

// ❌ SEM ponteiro - CRIA UMA CÓPIA
func (c ContaCorrente) SacarComCopia(valor float64) {
    fmt.Printf("Dentro da função - Endereço: %p\n", &c)
    c.Saldo -= valor  // Modifica apenas a CÓPIA local
    fmt.Printf("Saldo na cópia: %.2f\n", c.Saldo)
    // Quando sai da função, a cópia é destruída
}

// ✅ COM ponteiro - USA O ORIGINAL  
func (c *ContaCorrente) SacarComPonteiro(valor float64) {
    fmt.Printf("Dentro da função - Endereço: %p\n", c)
    c.Saldo -= valor  // Modifica a struct ORIGINAL
    fmt.Printf("Saldo original: %.2f\n", c.Saldo)
}

func main() {
    conta := ContaCorrente{Saldo: 1000}
    fmt.Printf("Conta original - Endereço: %p\n", &conta)
    
    // Teste com cópia
    fmt.Println("\n=== TESTANDO COM CÓPIA ===")
    conta.SacarComCopia(100)
    fmt.Printf("Saldo após saque com cópia: %.2f\n", conta.Saldo)  // Ainda 1000!
    
    // Teste com ponteiro
    fmt.Println("\n=== TESTANDO COM PONTEIRO ===")
    conta.SacarComPonteiro(100)  
    fmt.Printf("Saldo após saque com ponteiro: %.2f\n", conta.Saldo)  // Agora 900!
}

/* Saída esperada:
Conta original - Endereço: 0xc000014078

=== TESTANDO COM CÓPIA ===
Dentro da função - Endereço: 0xc000014088  // Endereço DIFERENTE!
Saldo na cópia: 900.00
Saldo após saque com cópia: 1000.00        // Original não mudou!

=== TESTANDO COM PONTEIRO ===  
Dentro da função - Endereço: 0xc000014078  // Mesmo endereço!
Saldo original: 900.00
Saldo após saque com ponteiro: 900.00      // Original foi modificado!
*/
```

## Análise dos Métodos

### 1. Método Sacar

```go
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
    podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
    if podeSacar {
        c.Saldo -= valorDoSaque  // Modifica o saldo ORIGINAL
        return "Saque realizado com sucesso"
    } else {
        return "Saldo insuficiente"
    }
}
```

**Ordem dos Parâmetros:**
1. **Receptor**: `(c *ContaCorrente)` - a struct que "recebe" o método (PONTEIRO)
2. **Parâmetro**: `valorDoSaque float64` - o valor a ser sacado
3. **Retorno**: `string` - mensagem de status

**Lógica:**
- Verifica se o valor é positivo E menor/igual ao saldo
- Se sim: deduz do saldo **original** e retorna sucesso
- Se não: retorna erro

### 2. Método Depositar

```go
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
    if valorDoDeposito > 0 {
        c.Saldo += valorDoDeposito  // Modifica o saldo ORIGINAL
        return "Deposito realizado com sucesso", c.Saldo
    } else {
        return "Valor do deposito menor que zero", c.Saldo
    }
}
```

**Ordem dos Parâmetros:**
1. **Receptor**: `(c *ContaCorrente)` - a struct que "recebe" o método (PONTEIRO)
2. **Parâmetro**: `valorDoDeposito float64` - o valor a ser depositado
3. **Retorno**: `(string, float64)` - **múltiplos retornos** (mensagem + saldo)

**Análise Detalhada do IF:**

A lógica condicional do método é baseada em uma validação simples mas importante:

```go
if valorDoDeposito > 0 {
    // FLUXO POSITIVO: Depósito válido
    c.Saldo += valorDoDeposito
    return "Deposito realizado com sucesso", c.Saldo
} else {
    // FLUXO NEGATIVO: Depósito inválido
    return "Valor do deposito menor que zero", c.Saldo
}
```

**Condição: `valorDoDeposito > 0`**
- **Propósito**: Impedir depósitos com valores negativos ou zero
- **Regra de negócio**: Só é possível depositar valores positivos
- **Segurança**: Evita que o saldo seja reduzido através de um "depósito negativo"

**Fluxo VERDADEIRO (valorDoDeposito > 0):**
```go
c.Saldo += valorDoDeposito  // Operador de soma e atribuição
return "Deposito realizado com sucesso", c.Saldo
```
- **`c.Saldo += valorDoDeposito`**: Equivale a `c.Saldo = c.Saldo + valorDoDeposito`
- **Modifica o original**: Como `c` é um ponteiro, altera a conta original na memória
- **Retorna**: Mensagem de sucesso + novo saldo atualizado

**Fluxo FALSO (valorDoDeposito <= 0):**
```go
return "Valor do deposito menor que zero", c.Saldo
```
- **Não modifica o saldo**: O saldo permanece inalterado
- **Retorna**: Mensagem de erro + saldo atual (sem alteração)
- **Inclui zero**: A condição `> 0` rejeita tanto valores negativos quanto zero

**Exemplos Práticos da Lógica:**

```go
conta := ContaCorrente{Saldo: 1000.0}

// ✅ Caso VÁLIDO - valor positivo
msg1, saldo1 := conta.Depositar(250.0)
// Resultado: msg1 = "Deposito realizado com sucesso", saldo1 = 1250.0
// conta.Saldo agora é 1250.0

// ❌ Caso INVÁLIDO - valor negativo  
msg2, saldo2 := conta.Depositar(-100.0)
// Resultado: msg2 = "Valor do deposito menor que zero", saldo2 = 1250.0
// conta.Saldo permanece 1250.0 (inalterado)

// ❌ Caso INVÁLIDO - valor zero
msg3, saldo3 := conta.Depositar(0.0)
// Resultado: msg3 = "Valor do deposito menor que zero", saldo3 = 1250.0
// conta.Saldo permanece 1250.0 (inalterado)
```

**Por que rejeitar zero?**
- **Operação sem sentido**: Depositar zero não altera o saldo
- **Evita processamento desnecessário**: Não há razão para processar uma operação nula
- **Clareza de intenção**: Força o usuário a especificar um valor real para depósito

**Características Especiais:**
- **Múltiplos retornos**: Go permite retornar mais de um valor simultaneamente
- **Retorna sempre o saldo**: Tanto em caso de sucesso quanto de erro, o saldo atual é retornado
- **Mensagem descritiva**: O retorno string informa claramente o resultado da operação
- **Saldo atualizado**: Em caso de sucesso, retorna o saldo já com o valor depositado

**Padrão de Design Observado:**
```go
// Padrão: (StatusMessage, UpdatedValue)
return "Mensagem de status", c.Saldo
```
Este padrão é útil porque:
- O caller pode verificar se a operação foi bem-sucedida pela mensagem
- O caller recebe imediatamente o valor atualizado sem precisar consultar novamente
- Facilita o debugging e logging de operações
````