# Módulo 3

## Conceito central

* **Interfaces** em Go são um conceito fundamental para o design de software.
* Elas permitem definir comportamentos abstratos que podem ser implementados por diferentes tipos, promovendo a flexibilidade e a composição em vez da herança tradicional.
* Neste módulo, vamos explorar como as interfaces funcionam em Go, como usá-las efetivamente e os padrões comuns relacionados a elas.

São muito diferentes de outras linguagens.

**Typescript**
```typescript
interface Notifier {
  send(message: string): vid;
}

// Declara explicitamente
class EmailNotifier implements Notifier {
  send(message: string) {
    // Enviar email
  }
}
```
**Go**
```go
type Notifierier interface {
  Send(message string) error
}

type EmailNotifier struct {
  // campos específicos do email
}

func (e EmailNotifier) Send(message string) error {
  // Enviar email
  return nil
}

// ✅ EmailNotifier implementa Notifier AUTOMATICAMENTE!
// Não há "implements", é IMPLÍCITO
```

Regra: **"Se tem os methods, implementa a interface."**

Não é necessário declarar que se implementa uma interface. Se a struct tem os methods corretos, automaticamente satisfaz a interface.

#### Interfaces básicas

Interface é um conjunto de assinaturas de métodos. Qualquer tipo que tenha esses métodos implementados satisfaz a interface.

```go
type Notifier interface {
    Send(message string) error
}

type Logger interface {
    Log(level string, message string)
}
```
**Convenção:** Nomes terminam em -er (Notifier, Logger, Reader, Writer)

### Implementação Implícita

```go
type EmailNotifier struct {
    SMTPServer string
}

// Se este method existe, EmailNotifier implementa Notifier!
func (e EmailNotifier) Send(message string) error {
    fmt.Printf("Sending email: %s via %s\n", message, e.SMTPServer)
    return nil
}

type SMSNotifier struct {
    PhoneNumber string
}

func (s SMSNotifier) Send(message string) error {
    fmt.Printf("Sending SMS to %s: %s\n", s.PhoneNumber, message)
    return nil
}
```

### Uso

```go
func Notify(n Notifier, message string) error {
    return n.Send(message)  // polimorfismo!
}

func main() {
    email := EmailNotifier{SMTPServer: "smtp.gmail.com"}
    sms := SMSNotifier{PhoneNumber: "+351912345678"}

    Notify(email, "Hello via email")  // ✅
    Notify(sms, "Hello via SMS")      // ✅

    // Ambos implementam Notifier!
}
```

## Por que de interfaces?

### 1. Polimorfismo

```go
type Player interface {
    GetStats() string
}

type FootballPlayer struct {
    Name  string
    Goals int
}

func (f FootballPlayer) GetStats() string {
    return fmt.Sprintf("%s: %d golos", f.Name, f.Goals)
}

type BasketballPlayer struct {
    Name   string
    Points int
}

func (b BasketballPlayer) GetStats() string {
    return fmt.Sprintf("%s: %d pontos", b.Name, b.Points)
}

func PrintStats(players []Player) {
    for _, p := range players {
        fmt.Println(p.GetStats())  // cada um implementa à sua maneira
    }
}
```

### 2. Dependency Injection (crucial para APIs)

```go
// Repository interface
type PlayerRepository interface {
    Save(player Player) error
    FindByID(id int) (Player, error)
}

// Implementação real (PostgreSQL)
type PostgresPlayerRepo struct {
    db *sql.DB
}

func (r PostgresPlayerRepo) Save(player Player) error {
    // INSERT INTO players ...
}

func (r PostgresPlayerRepo) FindByID(id int) (Player, error) {
    // SELECT * FROM players WHERE id = ...
}

// Implementação mock (testes)
type MockPlayerRepo struct {
    players map[int]Player
}

func (r MockPlayerRepo) Save(player Player) error {
    r.players[player.ID] = player
    return nil
}

func (r MockPlayerRepo) FindByID(id int) (Player, error) {
    return r.players[id], nil
}

// Service não sabe qual implementação está a usar!
type PlayerService struct {
    repo PlayerRepository  // interface!
}

func (s PlayerService) CreatePlayer(p Player) error {
    return s.repo.Save(p)  // funciona com qualquer implementação
}
```

### Vantagens:

* ✅ Testes fáceis (usa mock)
* ✅ Troca implementação sem mudar código
* ✅ Desacopla camadas

### 3. Interfaces vazias (`interface{}`)

```go
// Go 1.18+
var x any  // pode ser QUALQUER coisa

// Equivalente a:
var x interface{}  // forma antiga

// Uso

func Print(value any) {
    fmt.Println(value)
}

Print(42)
Print("hello")
Print(Player{Name: "Ronaldo"})
```

### Parte 4: Type Assertions

```go
var i interface{} = "hello"

// Type assertion
s := i.(string)  // ✅ "hello"
fmt.Println(s)

// n := i.(int)  // ❌ PANIC! i não é int

// Safe assertion (com check)
n, ok := i.(int)
if ok {
    fmt.Println("É int:", n)
} else {
    fmt.Println("Não é int")  // ← entra aqui
}
```

### Parte 5: Type Switch

```go
func Describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("É um int: %d\n", v)
    case string:
        fmt.Printf("É uma string: %s\n", v)
    case Player:
        fmt.Printf("É um Player: %s\n", v.Name)
    default:
        fmt.Printf("Tipo desconhecido: %T\n", v)
    }
}

Describe(42)                    // É um int: 42
Describe("hello")               // É uma string: hello
Describe(Player{Name: "Messi"}) // É um Player: Messi
```

### Parte 6: Interface Composition

```go
type Reader interface {
    Read() ([]byte, error)
}

type Writer interface {
    Write(data []byte) error
}

// Composição
type ReadWriter interface {
    Reader  // tem Read()
    Writer  // tem Write()
}

// Para implementar ReadWriter, precisa de Read() E Write()
type File struct{}

func (f File) Read() ([]byte, error) {
    // ...
}

func (f File) Write(data []byte) error {
    // ...
}

// File implementa ReadWriter automaticamente!
```

### Parte 7: Interfaces da Standard Library

```go
// fmt.Stringer - custom string representation
type Stringer interface {
    String() string
}

type Player struct {
    Name  string
    Goals int
}

func (p Player) String() string {
    return fmt.Sprintf("%s (%d golos)", p.Name, p.Goals)
}

func main() {
    p := Player{Name: "Ronaldo", Goals: 25}
    fmt.Println(p)  // chama p.String() automaticamente!
    // Output: Ronaldo (25 golos)
}
```

```go
// error - já conheces!
type error interface {
    Error() string
}

// Podes criar erros custom
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidationError implementa error automaticamente!
```

### Problema pointer receiver vs value receiver

```go
type Notifier interface {
    Send(msg string) error
}

type Email struct{}

func (e *Email) Send(msg string) error {  // ← POINTER receiver
    return nil
}

func main() {
    var n Notifier = Email{}  // ❌ ERRO!
}
```

**Erro de compilação**:
```bash
Email does not implement Notifier (Send method has pointer receiver)
```

Por quê?
Em Go:

* Se o method tem pointer receiver `*Email`
* Apenas `*Email` implementa a interface
* `Email` (value) NÃO implementa

Regra
```go
func (e *Email) Send() error  // só *Email implementa interface
func (e Email) Send() error   // Email e *Email implementam interface
```

Soluções
```go
// Solução 1: Usar ponteiro (mais comum)
var n Notifier = &Email{}  // ✅ OK

// Solução 2: Mudar para value receiver
func (e Email) Send(msg string) error {  // sem *
    return nil
}
var n Notifier = Email{}  // ✅ OK
```

|Method Receiver|`*Email{}` implementa? |`&Email{}` implementa?|
|-----------------|--------------------|--------------------|
|`(e Email)`      | ✅ SIM             | ✅ SIM              |
|`(e *Email)`     | ❌ NÃO             | ✅ SIM              |

Resumo:
* Value receiver é mais flexível (funciona com ambos).
* Pointer receiver é necessário se o method modifica o struct ou se a struct é grande.

🎓 Resumo - Quando Usar Cada Um
* Value Receiver (t Type)
* ✅ Usa quando:
  * Method só lê dados (não modifica)
  * Struct é pequena (poucos campos, tipos primitivos)
* Quer que value e pointer implementem interface

* ❌ NÃO usa quando:
  * Method modifica a struct
  * Struct é grande (cópia custosa)
  * Struct tem mutex ou channels

Memorizar:

* Pointer receiver → Só &Type{} implementa interface
* Value receiver → Ambos implementam interface

💡 Por Quê Esta Diferença?

Chamar method diretamente:
```go
f := File{Name: "test"}
f.Read()  // Go pode fazer &f automaticamente (sabe o endereço)
```
Go sabe onde está `f` na memória, pode criar ponteiro.

Atribuir a interface:
```go
var r Reader = File{Name: "test"}  // ❌
```

* Quando crias `File{}`, é um valor temporário. Go não consegue criar ponteiro para algo temporário que vai desaparecer!
* Tecnicamente: Interfaces armazenam um ponteiro para o valor. Se o method precisa de `*File`, a interface precisa de `*File`.

Value receiver → flexível
```go
func (e Email) Send(msg string) error { ... }

var n1 Notifier = Email{}   // ✅
var n2 Notifier = &Email{}  // ✅
```

 Pointer receiver → só ponteiro
 ```go
 func (e *Email) Send(msg string) error { ... }

var n1 Notifier = Email{}   // ❌
var n2 Notifier = &Email{}  // ✅
 ```