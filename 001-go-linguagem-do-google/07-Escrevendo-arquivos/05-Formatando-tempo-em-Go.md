# 05 - Formatando tempo em Go

---

## Transcrição

Queremos imprimir a data no log no seguinte formato:

```
05/06/2017 10:50:06 - https://www.alura.com.br - online: true
```

Para trabalhar com tempo em Go, vamos utilizar o já conhecido pacote `time`. Para pegar o tempo atual, ele possui a função `Now`, que nos retorna um objeto, portanto vamos formatá-lo, com a função `format`.

Para formatar a data, essa função utiliza diversas constantes, que podem ser consultadas no seu código fonte:

```go
const (
      _                        = iota
      stdLongMonth             = iota + stdNeedDate  // "January"
      stdMonth                                       // "Jan"
      stdNumMonth                                    // "1"
      stdZeroMonth                                   // "01"
      stdLongWeekDay                                 // "Monday"
      stdWeekDay                                     // "Mon"
      stdDay                                         // "2"
      stdUnderDay                                    // "_2"
      stdZeroDay                                     // "02"
      stdHour                  = iota + stdNeedClock // "15"
      stdHour12                                      // "3"
      stdZeroHour12                                  // "03"
      stdMinute                                      // "4"
      stdZeroMinute                                  // "04"
      stdSecond                                      // "5"
      stdZeroSecond                                  // "05"
      stdLongYear              = iota + stdNeedDate  // "2006"
      stdYear                                        // "06"
      stdPM                    = iota + stdNeedClock // "PM"
      stdpm                                          // "pm"
      stdTZ                    = iota                // "MST"
      stdISO8601TZ                                   // "Z0700"  // prints Z for UTC
      stdISO8601SecondsTZ                            // "Z070000"
      stdISO8601ShortTZ                              // "Z07"
      stdISO8601ColonTZ                              // "Z07:00" // prints Z for UTC
      stdISO8601ColonSecondsTZ                       // "Z07:00:00"
      stdNumTZ                                       // "-0700"  // always numeric
      stdNumSecondsTz                                // "-070000"
      stdNumShortTZ                                  // "-07"    // always numeric
      stdNumColonTZ                                  // "-07:00" // always numeric
      stdNumColonSecondsTZ                           // "-07:00:00"
      stdFracSecond0                                 // ".0", ".00", ... , trailing zeros included
      stdFracSecond9                                 // ".9", ".99", ..., trailing zeros omitted

      stdNeedDate  = 1 << 8             // need month, day, year
      stdNeedClock = 2 << 8             // need hour, minute, second
      stdArgShift  = 16                 // extra argument in high bits, above low stdArgShift
      stdMask      = 1<<stdArgShift - 1 // mask out argument
  )
```

Como queremos representar dia, mês e ano no formato numérico, vamos utilizar o `stdZeroDay`, representado pelo `02`, `stdZeroMonth`, representado pelo `01` e `stdLongYear`, representado pelo `2006`:

```go
// restante do código omitido

func registraLog(site string, status bool) {

    arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    arquivo.WriteString(time.Now().Format("02/01/2006") + " - " + site + " - online: " +
        strconv.FormatBool(status) + "\n")

    arquivo.Close()
}
```

Falta ajustar a hora, vamos utilizar o `stdHour`, representado pelo `15`, `stdZeroMinute`, representado pelo `04`, e `stdZeroSecond`, representado pelo `05`:

```go
// restante do código omitido

func registraLog(site string, status bool) {

    arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + 
        " - online: " + strconv.FormatBool(status) + "\n")

    arquivo.Close()
}
```

Podemos agora executar o monitoramento e verificar o log, ele está exatamente como queríamos! Agora que temos a funcionalidade de salvar os logs pronta, falta exibi-los ao escolher o comando 2, e é justamente isso que faremos no próximo vídeo. 




## Testando

- Testando a geração de log com tempo registrado:


> cd /home/fernando/cursos/golang/formacao-go-alura/001-go-linguagem-do-google/07-Escrevendo-arquivos
> go run 05-formatando-tempo.go
Olá, sr. Douglas
Este programa está na versão 1.2
&{0xc00008c240}
1- Iniciar Monitoramento
2- Exibir Logs
0- Sair do Programa
1
O comando escolhido foi 1

Monitorando...
Testando site 0 : https://www.google.com
Site: https://www.google.com foi carregado com sucesso!
&{0xc000142000}
Testando site 1 : https://www.github.com
Site: https://www.github.com foi carregado com sucesso!
&{0xc000250ae0}


- Arquivo de log:

https://www.github.com - online: true
https://www.alura.com.br - online: true
https://www.stackoverflow.com - online: false
https://www.golang.org - online: true
04/10/2025 23:31:13 - teste-site-falso - online: false
04/10/2025 23:31:16 - https://www.google.com - online: true
04/10/2025 23:31:16 - https://www.github.com - online: true
04/10/2025 23:31:17 - https://www.alura.com.br - online: true
04/10/2025 23:31:17 - https://www.stackoverflow.com - online: false
04/10/2025 23:31:19 - https://www.golang.org - online: true
04/10/2025 23:31:24 - https://www.google.com - online: true
04/10/2025 23:31:24 - https://www.github.com - online: true
04/10/2025 23:31:24 - https://www.alura.com.br - online: true
04/10/2025 23:31:24 - https://www.stackoverflow.com - online: false
04/10/2025 23:31:25 - https://www.golang.org - online: true



O que você está vendo &{0xc000142000} é a referência de memória do objeto arquivo sendo impressa no terminal.

Isso acontece porque você tem esta linha na função registraLog():

O que estava acontecendo:

fmt.Println(arquivo) imprime a estrutura interna do objeto arquivo
&{0xc000142000} é o endereço de memória onde o arquivo está armazenado
Isso é útil apenas para debug, não para o usuário final


- Ajustando arquivo Go.

- Testando, não trouxe mais a referencia de memória do objeto:

> go run 05-formatando-tempo.go
Olá, sr. Douglas
Este programa está na versão 1.2
1- Iniciar Monitoramento
2- Exibir Logs
0- Sair do Programa
1
O comando escolhido foi 1

Monitorando...
Testando site 0 : https://www.google.com
Site: https://www.google.com foi carregado com sucesso!
Testando site 1 : https://www.github.com
Site: https://www.github.com foi carregado com sucesso!
Testando site 2 : https://www.alura.com.br

- Logs gerados:

04/10/2025 23:36:13 - https://www.google.com - online: true
04/10/2025 23:36:13 - https://www.github.com - online: true
04/10/2025 23:36:14 - https://www.alura.com.br - online: true
04/10/2025 23:36:14 - https://www.stackoverflow.com - online: false




## RESUMO
Nesta aula do curso "Go: a linguagem do Google", aprendemos a trabalhar com o pacote time para formatar e registrar a data e hora em um log. 
O objetivo é imprimir a data no formato específico: 05/06/2017 10:50:06 - https://www.alura.com.br - online: true.

Para isso, utilizamos a função Now() para obter o tempo atual e a função Format() para formatá-lo. 

Usamos constantes do pacote time para representar o dia, mês, ano e hora no formato desejado. 
Go usa um sistema único de formatação baseado em uma data de referência específica:
| Elemento | Constante | Valor de Referência |
|----------|-----------|-------------------|
| Dia | `stdZeroDay`     | `02` |
| Mês | `stdZeroMonth`   | `01` |
| Ano | `stdLongYear`    | `2006` |
| Hora | `stdHour`      | `15` |
| Minuto | `stdZeroMinute`   | `04` |
| Segundo | `stdZeroSecond`  | `05` |
Na documentação do Go temos os detalhes:
<https://go.dev/src/time/format.go>
Esses nomes stdZeroMonth, stdLongYear não são para usar diretamente; eles só ajudam a entender por que 01 adiciona zero e 1 não, ou por que 2006 é o ano “longo”.

Para BR no dia a dia: use "02/01/2006 15:04:05" ao exibir e time.RFC3339 em logs/integrações.

No código, registramos as informações em um arquivo chamado "log.txt", incluindo a data e hora formatadas, o site e o status de online.
A aula também abordou como abrir o arquivo para escrita e garantir que os logs sejam salvos corretamente. 
A próxima etapa será exibir esses logs ao escolher um comando específico.