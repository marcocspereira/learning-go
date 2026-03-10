# 💡 Conceito em 2 Minutos

**Analogia:** Errors em Go são como códigos de retorno específicos num restaurante:

```text
Cliente pede prato → Cozinha pode retornar:
- nil (tudo OK, aqui está o prato)
- ErrOutOfStock (não temos ingredientes)
- ErrKitchenClosed (cozinha fechada)
- ErrInvalidOrder (pedido inválido)
```

Cliente verifica o código antes de comer. Se for erro, sabe exatamente o que aconteceu.

## 🔑 3 Técnicas Essenciais

### 1. Sentinel Errors (Erros Pré-definidos)

```go
// Define erros standard
var (
    ErrNotFound      = errors.New("not found")
    ErrAlreadyExists = errors.New("already exists")
    ErrInvalidData   = errors.New("invalid data")
)

// Usa
func (repo *InMemoryPlayerRepo) FindByID(id int) (Player, error) {
    player, exists := repo.players[id]
    if !exists {
        return Player{}, ErrNotFound  // ← erro fixo
    }
    return player, nil
}

// Quem chama pode verificar
player, err := repo.FindByID(123)
if errors.Is(err, ErrNotFound) {  // ← verifica tipo específico
    fmt.Println("Jogador não existe")
}
```

* **Vantagem:** Podes identificar tipo de erro, não só mensagem.
* **Usar quando:**
  * ✅ O mesmo tipo de erro acontece em vários sítios
  * ✅ Quem chama precisa reagir diferente para cada tipo
  * ✅ Erro é "standard" (not found, already exists, permission denied)
* **Quando NÃO usar:** Erros únicos que só acontecem num sítio.

```go
// Define uma vez
var ErrNotFound = errors.New("not found")

// Usa em vários repos
func (r *PlayerRepo) FindByID(id int) (Player, error) {
    if !exists {
        return Player{}, ErrNotFound  // ✅
    }
}

func (r *TeamRepo) FindByID(id int) (Team, error) {
    if !exists {
        return Team{}, ErrNotFound  // ✅ mesmo erro
    }
}

// Quem chama pode tratar de forma especial
player, err := repo.FindByID(123)
if errors.Is(err, ErrNotFound) {
    // Retorna HTTP 404
} else if err != nil {
    // Retorna HTTP 500
}
```

### 2. Custom Error Types (Erros com Contexto)

```go
// Erro com dados adicionais
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// Usa
func (p Player) Validate() error {
    if p.Number < 1 || p.Number > 99 {
        return ValidationError{
            Field:   "Number",
            Message: "must be between 1 and 99",
        }
    }
    return nil
}

// Quem chama pode extrair dados
err := player.Validate()
if verr, ok := err.(ValidationError); ok {
    fmt.Printf("Campo %s inválido: %s\n", verr.Field, verr.Message)
}
```

* **Vantagem:** Erro carrega informação adicional.
* **Quando usar:**
  * ✅ Precisas guardar dados no erro (campo, valor, código)
  * ✅ Quem chama precisa extrair informação do erro
  * ✅ Erro tem contexto estruturado
* **Quando NÃO usar:** Erro simples sem contexto adicional.

```go
type ValidationError struct {
    Field   string
    Value   interface{}
    Message string
}

func (p Player) Validate() error {
    if p.Number < 1 || p.Number > 99 {
        return ValidationError{
            Field:   "Number",
            Value:   p.Number,      // ← guarda o valor inválido!
            Message: "must be 1-99",
        }
    }
}

// API pode usar os dados
err := player.Validate()
if verr, ok := err.(ValidationError); ok {
    return JSONError{
        Field:   verr.Field,    // ← extrai dados
        Message: verr.Message,
        Value:   verr.Value,
    }
}
```

### 3. Error Wrapping `fmt.Errorf`

```go
// Adiciona contexto SEM perder erro original
func (s *PlayerService) CreatePlayer(player Player) (Player, error) {
    _, err := s.teamRepo.FindByID(player.TeamID)
    if err != nil {
        return Player{}, fmt.Errorf("failed to validate team: %w", err)
        //                                                      ↑ %w preserva erro
    }
    return s.playerRepo.Create(player)
}

// Stack de erros:
// "failed to validate team: team not found"
```

* **Vantagem:** Preserva erro original, adiciona contexto.
* **Quando usar:**
  * ✅ Erro é único/específico a uma situação
  * ✅ Não precisas identificar o tipo depois
  * ✅ Só queres mensagem descritiva
```go
func ProcessMatch(homeScore, awayScore int) error {
    if homeScore < 0 {
        return fmt.Errorf("home score cannot be negative: %d", homeScore)
    }
    // Ninguém vai verificar "se é erro de score negativo"
    // Só vai logar a mensagem
}
```

## 🔑 Quando Usar Cada Tipo de Erro

📊 Decision Tree Simples
```text
Preciso de informação extra sobre o erro?
│
├─ NÃO, só preciso saber "que tipo" de erro
│  │
│  └─ É erro que acontece em MUITOS sítios?
│     ├─ SIM → Sentinel Error (var ErrNotFound)
│     └─ NÃO → fmt.Errorf normal
│
└─ SIM, preciso de dados (campo, valor, etc)
   └─ Custom Error Type (struct ValidationError)
```

---

## 💥 Panic - O "Emergency Stop"

💡 Analogia
**Error** = Semáforo vermelho
* Paras o carro
* Verificas se podes continuar
* Decides o que fazer

**Panic** = Travão de emergência do comboio

*  Para TUDO imediatamente
*  Não podes continuar normalmente
*  Algo impossível aconteceu

### 🔴 O Que é Panic?
**Panic** = crash controlado do programa.

```go
func divide(a, b int) int {
    if b == 0 {
        panic("division by zero!")  // ← PARA TUDO!
    }
    return a / b
}

func main() {
    result := divide(10, 0)  // ← programa CRASH aqui
    fmt.Println(result)      // ← NUNCA chega aqui
}

// Output:
// panic: division by zero!
// goroutine 1 [running]:
// main.divide(...)
// ...
```

#### Quando faz panic:

* ❌ Para execução da função imediatamente
* ❌ Volta atrás na stack (unwind)
* ❌ Programa termina (a não ser que uses recover)

#### 🎯 Quando Usar Panic?
NUNCA usar **panic** para:
* ❌ Erros normais (user input inválido, file not found)
* ❌ Erros esperados
* ❌ Condições de negócio

SÓ usa **panic** para:
* ✅ Bugs no teu código (programação defensiva)
* ✅ Situações impossíveis (nunca deveria acontecer)
* ✅ Inicialização falhada (config inválido ao arrancar)

### Exemplos práticos

#### ❌ NÃO uses panic:
```go
// ❌ ERRADO - user input é erro normal
func CreatePlayer(name string) Player {
    if name == "" {
        panic("name cannot be empty")  // ❌ NÃO!
    }
}

// ✅ CORRETO - retorna erro
func CreatePlayer(name string) (Player, error) {
    if name == "" {
        return Player{}, errors.New("name cannot be empty")
    }
}
```

#### ✅ USA panic:

```go
// ✅ OK - bug no código (impossível)
func (repo *InMemoryPlayerRepo) mustGetNextID() int {
    if repo.nextID < 0 {
        panic("nextID is negative - impossible state")
    }
    return repo.nextID
}

// ✅ OK - configuração inválida (ao arrancar)
func main() {
    config := loadConfig()
    if config.DatabaseURL == "" {
        panic("DATABASE_URL not set")  // ✅ não pode continuar
    }
    // ...
}

// ✅ OK - índice out of bounds (bug)
func getPlayer(players []Player, index int) Player {
    if index < 0 || index >= len(players) {
        panic(fmt.Sprintf("index %d out of bounds", index))
    }
    return players[index]
}
```

## 🛡️ Recover (Apanhar Panic)
Podes "apanhar" um panic com recover():
```go
func safeOperation() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Apanhei panic:", r)
        }
    }()

    panic("something went wrong")  // ← panic aqui
    fmt.Println("nunca executa")
}

func main() {
    safeOperation()
    fmt.Println("Programa continua!")  // ← executa!
}

// Output:
// Apanhei panic: something went wrong
// Programa continua!
```

Usado em: HTTP handlers (não queres que 1 request crashe o servidor)
```go
func handler(w http.ResponseWriter, r *http.Request) {
    defer func() {
        if r := recover(); r != nil {
            http.Error(w, "Internal Server Error", 500)
            log.Printf("Panic: %v", r)
        }
    }()

    // código que pode fazer panic
}
```

### 🎯 Regra de Ouro
```go
// ✅ 99% do tempo - usa errors
func DoSomething() error {
    if problem {
        return errors.New("problem occurred")
    }
    return nil
}

// ❌ Raramente - usa panic
func DoSomething() {
    if impossibleCondition {
        panic("this should never happen")
    }
}
```

Em Go idiomático: Erros são valores, não exceptions. Trata-os explicitamente.