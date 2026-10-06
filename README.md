# 📖 Estudo: A Linguagem de Programação Go

Repositório dedicado ao estudo aprofundado do livro **"A Linguagem de Programação Go"** (*The Go Programming Language*), de **Alan A. A. Donovan** e **Brian W. Kernighan**.

O objetivo deste projeto é registrar anotações teóricas, implementar os exemplos de cada capítulo e resolver todos os exercícios práticos propostos pelos autores.

---

## 🎯 Progresso Atual (Capítulo 1: Tutorial)

| Exercício / Código | Descrição | Status |
| :--- | :--- | :---: |
| [`ch01/01`](ch01/01/main.go) | **Ex. 1.1**: Modificar `echo` para exibir também `os.Args[0]` (nome do comando). | Concluído ✅ |
| [`ch01/02`](ch01/02/main.go) | **Ex. 1.2**: Modificar `echo` para exibir o índice e o valor de cada argumento por linha. | Concluído ✅ |
| [`ch01/03`](ch01/03/main.go) | **Ex. 1.3**: Medir a diferença de desempenho entre concatenação com `+` vs `strings.Join`. | Concluído ✅ |
| [`ch01/04`](ch01/04/main.go) | **Ex. 1.4**: Modificar `dup2` para imprimir os nomes de todos os arquivos onde linhas duplicadas aparecem. | Concluído ✅ |
| [`ch01/05`](ch01/05/main.go) | **Ex. 1.5**: Alterar paleta de Lissajous para traçado verde sobre fundo preto. | Concluído ✅ |
| [`ch01/06`](ch01/06/main.go) | **Ex. 1.6**: Alterar gerador de Lissajous para produzir figuras em múltiplas cores. | Concluído ✅ |
| [`ch01/07`](ch01/07/main.go) | **Ex. 1.7**: Usar `io.Copy(os.Stdout, resp.Body)` em `fetch` em vez de `ioutil.ReadAll`. | Concluído ✅ |
| [`ch01/08`](ch01/08/main.go) | **Ex. 1.8**: Modificar `fetch` para adicionar prefixo `http://` caso omitido. | Concluído ✅ |
| [`ch01/09`](ch01/09/main.go) | **Ex. 1.9**: Modificar `fetch` para imprimir o status HTTP retornado (`resp.StatusCode`). | Concluído ✅ |
| [`ch01/lissajous.go`](ch01/lissajous.go) | Implementação base da animação GIF de figuras de Lissajous. | Concluído ✅ |

---

## 📅 Plano de Estudos Semanal (12 Semanas)

Cronograma estruturado para percorrer os 13 capítulos do livro com leitura, prática e exercícios:

- [x] **Semana 01 — Capítulo 1: Tutorial Inicial**
  - [x] Introdução, sintaxe básica e estrutura de comandos (`echo`, `os.Args`).
  - [x] Detecção de linhas duplicadas (`dup1`, `dup2`, `dup3`, `bufio`, `map`).
  - [x] Gráficos animados em GIF (figuras de Lissajous, pacotes `image`, `color`).
  - [x] Requisições HTTP com `net/http` e concorrência básica com `fetchall`.
  - [x] Servidor web simples e roteamento de rotas.
  - [x] Resolução dos exercícios 1.1 ao 1.9.

- [ ] **Semana 02 — Capítulo 2: Estrutura do Programa & Capítulo 3: Tipos de Dados Básicos**
  - [ ] Nomes, declarações, variáveis (`var`, `:=`), ponteiros, atribuições e tipos definidos pelo usuário.
  - [ ] Pacotes, importações, inicialização (`init()`) e escopo de identificadores.
  - [ ] Inteiros, pontos flutuantes, números complexos e booleanos.
  - [ ] Strings, literais de string, UTF-8, conversões e pacote `unicode/utf8`.
  - [ ] Constantes e gerador de enumeração `iota`.
  - [ ] Exercícios dos Capítulos 2 e 3.

- [ ] **Semana 03 — Capítulo 4: Tipos Compostos**
  - [ ] Arrays vs Slices: representação interna, fatiamento, `append` e `copy`.
  - [ ] Maps: tabelas hash em Go, ciclo de vida, testes de existência e ordenação de chaves.
  - [ ] Structs: campos, literais, ponteiros para structs e campos embutidos (*struct embedding*).
  - [ ] Serialização JSON: tags de struct, `json.Marshal`, `json.Unmarshal`.
  - [ ] Templates de texto e HTML (`text/template`, `html/template`).
  - [ ] Exercícios do Capítulo 4.

- [ ] **Semana 04 — Capítulo 5: Funções**
  - [ ] Declaração de funções, parâmetros e múltiplos retornos nomeados.
  - [ ] Tratamento idiomático de erros (estratégias de propagação e fim de arquivo).
  - [ ] Funções de primeira classe, valores de funções e closures anônimas.
  - [ ] Funções variádicas (`...interface{}`).
  - [ ] Gerenciamento de recursos com `defer`.
  - [ ] Pânico e recuperação (`panic`, `recover`).
  - [ ] Exercícios do Capítulo 5.

- [ ] **Semana 05 — Capítulo 6: Métodos & Orientação a Objetos em Go**
  - [ ] Declaração de métodos em tipos definidos pelo usuário.
  - [ ] Receptores de valor (*value receiver*) vs Receptores de ponteiro (*pointer receiver*).
  - [ ] Composição de tipos por incorporação de structs (*embedding*).
  - [ ] Valores de métodos e expressões de métodos.
  - [ ] Encapsulamento através de maiúsculas/minúsculas de identificadores.
  - [ ] Exercícios do Capítulo 6.

- [ ] **Semana 06 — Capítulo 7: Interfaces (Parte 1)**
  - [ ] Interfaces como contratos de comportamento e polimorfismo.
  - [ ] Satisfação implícita de interfaces.
  - [ ] Valores de interface, interface vazia (`any` / `interface{}`) e armadilhas com `nil`.
  - [ ] Interfaces padrão fundamentais: `io.Reader`, `io.Writer`, `fmt.Stringer`.
  - [ ] Exercícios iniciais do Capítulo 7.

- [ ] **Semana 07 — Capítulo 7: Interfaces (Parte 2)**
  - [ ] Ordenação com `sort.Interface`.
  - [ ] Servidores HTTP com `http.Handler` e multiplexadores.
  - [ ] A interface `error`.
  - [ ] Asserções de tipo (*type assertions*) e chaves de tipo (*type switches*).
  - [ ] Exercícios avançados do Capítulo 7.

- [ ] **Semana 08 — Capítulo 8: Goroutines e Canais (Concorrência Básica)**
  - [ ] Goroutines e modelo CSP (*Communicating Sequential Processes*).
  - [ ] Canais sem buffer e com buffer: sincronização e comunicação.
  - [ ] Pipelines, canais unidirecionais (`<-chan`, `chan<-`) e fechamento de canais.
  - [ ] Multiplexação com `select`, timeout e polling não-bloqueante.
  - [ ] Cancelamento gracioso de operações.
  - [ ] Exercícios do Capítulo 8.

- [ ] **Semana 09 — Capítulo 9: Concorrência com Variáveis Compartilhadas**
  - [ ] Condições de corrida (*race conditions*) e detector de corrida (`go test -race`).
  - [ ] Exclusão mútua com `sync.Mutex` e `sync.RWMutex`.
  - [ ] Inicialização sob demanda com `sync.Once`.
  - [ ] Sincronização de memória e pacote `sync/atomic`.
  - [ ] Comparação: Goroutines vs Threads do Sistema Operacional.
  - [ ] Exercícios do Capítulo 9.

- [ ] **Semana 10 — Capítulo 10: Pacotes e a Ferramenta Go & Capítulo 11: Testes**
  - [ ] Organização de módulos (`go mod`), ciclo de build e documentação com `godoc`.
  - [ ] Testes unitários com a ferramenta `go test`.
  - [ ] Testes orientados a tabela (*table-driven tests*).
  - [ ] Cobertura de código (`-cover`).
  - [ ] Testes de benchmark (`Benchmark*`, `b.N`) e exemplos executáveis (`Example*`).
  - [ ] Exercícios dos Capítulos 10 e 11.

- [ ] **Semana 11 — Capítulo 12: Reflexão (Reflection)**
  - [ ] Por que usar e quando evitar reflexão.
  - [ ] `reflect.Type` e `reflect.Value`.
  - [ ] Inspecionando structs, campos e tags em tempo de execução.
  - [ ] Modificando valores via reflexão.
  - [ ] Exercícios do Capítulo 12.

- [ ] **Semana 12 — Capítulo 13: Programação de Baixo Nível & Revisão Final**
  - [ ] Pacote `unsafe`: ponteiros brutos (`unsafe.Pointer`), tamanhos de memória (`unsafe.Sizeof`, `Alignof`, `Offsetof`).
  - [ ] Interoperabilidade com linguagem C via `cgo`.
  - [ ] Revisão geral dos conceitos fundamentais de Go e consolidação do repositório.

---

## 📁 Estrutura de Diretórios

```text
├── .vscode/               # Configurações de depuração e workspace do VS Code
├── ch01/                  # Capítulo 1: Tutorial
│   ├── 01/                # Ex. 1.1: os.Args[0]
│   ├── 02/                # Ex. 1.2: índice e valor de os.Args
│   ├── 03/                # Ex. 1.3: benchmark de concatenação vs strings.Join
│   ├── 04/                # Ex. 1.4: dup2 com listagem de arquivos
│   ├── 05/                # Ex. 1.5: Lissajous em verde e preto
│   ├── 06/                # Ex. 1.6: Lissajous multicor
│   ├── 07/                # Ex. 1.7: fetch usando io.Copy
│   ├── 08/                # Ex. 1.8: fetch com prefixo http:// automático
│   ├── 09/                # Ex. 1.9: fetch exibindo código de status HTTP
│   └── lissajous.go       # Exemplo base do gerador de GIFs de Lissajous
└── README.md              # Plano de estudos e acompanhamento
```

---

## 🚀 Como Executar

Certifique-se de ter o [Go instalado](https://go.dev/dl/) (versão 1.18 ou superior).

Para executar qualquer um dos exercícios, navegue até a pasta ou execute diretamente:

```bash
# Exemplo 1: Exibir argumentos e índices (Ex. 1.2)
go run ./ch01/02/main.go primeiro segundo terceiro

# Exemplo 2: Benchmark de concatenação vs strings.Join (Ex. 1.3)
go run ./ch01/03/main.go teste de desempenho com multiplas palavras

# Exemplo 3: Gerar animação de Lissajous em GIF (Ex. 1.5 / 1.6)
go run ./ch01/06/main.go > out.gif

# Exemplo 4: Fazer fetch de URL com status HTTP (Ex. 1.9)
go run ./ch01/09/main.go golang.org
```

---

## 📚 Referências

- **Livro:** Donovan, Alan A. A.; Kernighan, Brian W. *A Linguagem de Programação Go*. Novatec / Addison-Wesley.
- **Site Oficial Go:** [golang.org](https://golang.org) / [go.dev](https://go.dev)
- **Documentação Oficial:** [go.dev/doc](https://go.dev/doc)
