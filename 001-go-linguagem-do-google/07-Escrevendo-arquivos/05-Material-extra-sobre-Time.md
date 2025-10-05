
---

# Fusos e “relógios” (SP/UTC)

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() // horário local do SO
	sp, _ := time.LoadLocation("America/Sao_Paulo")

	fmt.Println("Agora (local do SO):", now)
	fmt.Println("Agora (São Paulo):  ", now.In(sp))
	fmt.Println("Agora (UTC):        ", now.UTC())
}
```

**Dica prática:** salve **em UTC** no banco e converta para “America/Sao\_Paulo” ao exibir.

---

# Formatação “BR” (números) e ISO para logs

Go usa a “data mágica” **2006-01-02 15:04:05** como layout.

```go
t := time.Date(2025, 10, 4, 14, 35, 0, 0, time.FixedZone("BRT", -3*3600))

fmt.Println(t.Format("02/01/2006 15:04:05")) // 04/10/2025 14:35:00  (BR)
fmt.Println(t.Format("2006-01-02 15:04:05")) // 2025-10-04 14:35:00
fmt.Println(t.Format(time.RFC3339))          // 2025-10-04T14:35:00-03:00 (bom p/ logs/APIs)
```

> Evite nomes de mês/dia (ex.: `Jan`, `Mon`) porque saem em inglês; prefira **numérico** no BR.

---

# Parse de string “BR” → `time.Time`

```go
s := "04/10/2025 14:35:00"
sp, _ := time.LoadLocation("America/Sao_Paulo")

t, err := time.ParseInLocation("02/01/2006 15:04:05", s, sp)
if err != nil { panic(err) }
fmt.Println(t) // 2025-10-04 14:35:00 -0300 -03
```

> Use **`ParseInLocation`** quando a string não traz fuso (ex.: `14:35:00`).
> Se vier ISO com fuso (`...-03:00`), use `time.Parse(time.RFC3339, s)`.

---

# Durações e prazos (expiração de link/QR)

```go
// expirar link em 15 minutos
criacao := time.Now()
ttl, _ := time.ParseDuration("15m")
expira := criacao.Add(ttl)

fmt.Println("Criação:", criacao.Format("02/01/2006 15:04:05"))
fmt.Println("Expira: ", expira.Format("02/01/2006 15:04:05"))
fmt.Println("Faltam: ", time.Until(expira)) // duration positiva ou negativa
```

---

# SLA/latência: medindo com precisão

```go
start := time.Now()
// ... chama serviço externo ...
elapsed := time.Since(start) // usa relógio monotônico (preciso p/ duração)
fmt.Println("Demorou:", elapsed)
```

---

# “Rodar de X em X tempo” (Ticker) e “esperar até” (Timer)

```go
// Ticker: executa a cada 30s dentro do horário comercial
ticker := time.NewTicker(30 * time.Second)
defer ticker.Stop()

sp, _ := time.LoadLocation("America/Sao_Paulo")
deadline := time.Now().Add(3 * time.Minute) // só de exemplo; poderia ser 18:00

for {
	select {
	case <-ticker.C:
		now := time.Now().In(sp)
		hh := now.Hour()
		if hh >= 9 && hh < 18 {
			fmt.Println("Coletando métricas às", now.Format("15:04:05"))
			// chama rotina...
		} else {
			fmt.Println("Fora do horário comercial:", now.Format("15:04:05"))
		}
	case <-time.After(time.Until(deadline)):
		fmt.Println("Parando o job periódico")
		return
	}
}
```

---

# HTTP com timeout (evitar travar)

```go
client := &http.Client{ Timeout: 5 * time.Second } // por requisição

resp, err := client.Get("https://api.seubanco.com/health")
if err != nil {
	fmt.Println("Erro/timeout:", err)
	return
}
defer resp.Body.Close()
fmt.Println("Status:", resp.StatusCode)
```

Para operações mais complexas/canceláveis, use **contexto**:

```go
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel()

req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://httpstat.us/200?sleep=4000", nil)
resp, err := http.DefaultClient.Do(req)
if err != nil {
	fmt.Println("Falhou (timeout?):", err)
	return
}
defer resp.Body.Close()
```

---

# “+3 dias úteis” (ex.: boleto) — simples (sem feriados)

```go
func AddBusinessDays(t time.Time, days int) time.Time {
	for days > 0 {
		t = t.Add(24 * time.Hour)
		if wd := t.Weekday(); wd != time.Saturday && wd != time.Sunday {
			days--
		}
	}
	return t
}

// uso:
sp, _ := time.LoadLocation("America/Sao_Paulo")
emissao := time.Date(2025, 10, 3, 10, 0, 0, 0, sp) // sex, 03/10/2025
venc := AddBusinessDays(emissao, 3)                // pula fim de semana
fmt.Println(venc.Format("02/01/2006"))             // 08/10/2025 (qua)
```

> Para **feriados**, mantenha uma lista (ou serviço) e cheque antes de contar o dia.

---

# Comparação de horários (janela de processamento)

```go
agora := time.Now()
inicioJanela, _ := time.ParseInLocation("15:04", "10:00", agora.Location())
fimJanela,   _ := time.ParseInLocation("15:04", "16:00", agora.Location())

// “cola” a data de hoje nesses horários
inicio := time.Date(agora.Year(), agora.Month(), agora.Day(),
	inicioJanela.Hour(), inicioJanela.Minute(), 0, 0, agora.Location())
fim := time.Date(agora.Year(), agora.Month(), agora.Day(),
	fimJanela.Hour(), fimJanela.Minute(), 0, 0, agora.Location())

if agora.Before(inicio) {
	fmt.Println("Ainda não abriu a janela")
} else if agora.After(fim) {
	fmt.Println("Janela fechada")
} else {
	fmt.Println("Dentro da janela")
}
```

---

# Boas práticas (resumo rápido)

* **Banco/integrações**: armazene em **UTC** (ex.: `TIMESTAMP WITH TIME ZONE`/ISO-8601). Converta para `America/Sao_Paulo` ao exibir.
* **Layouts**: use **`02/01/2006 15:04:05`** para BR e **`time.RFC3339`** para logs/APIs.
* **`ParseInLocation`**: quando a string **não** traz fuso.
* **Timeouts**: defina sempre em HTTP, DB, filas, etc. Prefira `context.WithTimeout` em fluxos longos.
* **Ticker/Timer**: sempre **pare** (`Stop()`) para não vazar goroutines.
* **Cálculo de duração**: use `time.Since`/`Sub` (monotônico).

---
 