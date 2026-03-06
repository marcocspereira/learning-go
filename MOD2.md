# Módulo 2

## Resumo
* Como funcionam ponteiros em Go
* Quando passar `*T` vs `T` para funções e métodos
* Passar por valor vs passar por referência
* Como structs e ponteiros interagem
* Gestão básica de memória

## Módulo 2.1: Ponteiros, Structs e Methods

### Ponteiros
* Guarda o endereço de memória de uma variável, não o valor em si.
* **Sintaxe:** `var p *int` declara um ponteiro para um inteiro.
* **Acessar o valor apontado:** `*p` (desreferenciar).
* **Atribuir um endereço:** `p = &x` (endereço de x).
* **Operadores:**
  * `&` = "endereço de" (address-of)
  * `*` = "valor em" (dereference)

```go
// Valor (cópia)
x := 42
fmt.Println("Valor de x:", x) // Valor de x: 42

// Ponteiro (referência)
y := &x // y &x guarda o endereço de x
fmt.Println("Endereço de x:", y) // Endereço de x: 0xc0000160a8 (endereço de memória)
fmt.Println("Valor apontado por y:", *y) // Valor apontado por y: 42

// Por valor (cópia)
func updateHeight(t Tree) {
    t.Height = 20  // modifica CÓPIA
}

// Por ponteiro (referência)
func updateHeight(t *Tree) {
    t.Height = 20  // modifica ORIGINAL
}
```

**Zero values:** Ponteiros não inicializados têm valor zero `nil`, indicando que não apontam para lugar nenhum. Deve verificar-se se um ponteiro é `nil` antes de desreferenciá-lo para evitar panics.

```go
var p *int
fmt.Println("Valor de p:", p) // Valor de p: <nil>
fmt.Println(*ptr) // ERRO: panic: runtime error: invalid memory address or nil pointer dereference
```

Exemplo:
```go
package main

import "fmt"

func duplicar(n int) {
    n = n * 2  // modifica cópia
}

func duplicarPtr(n *int) {
    *n = *n * 2  // modifica original
}

func main() {
    valor := 10

    duplicar(valor)
    fmt.Println(valor)  // 10 (não mudou)

    duplicarPtr(&valor)
    fmt.Println(valor)  // 20 (mudou!)
}
```

### Structs

* Structs são como "classes" sem herança. Agrupam dados relacionados.
* Passar struct por valor cria uma cópia. Passar por ponteiro permite modificar o original.

```go
// Definição
type Tree struct {
    ID       int
    Species  string
    Height   float64  // metros
    Diameter float64  // cm
}

// Criação
tree1 := Tree{
    ID:       1,
    Species:  "Pinus pinaster",
    Height:   15.5,
    Diameter: 25.3,
}

// Acesso
fmt.Println(tree1.Species)  // "Pinus pinaster"
tree1.Height = 16.0         // modificação
```

**Zero value de struct:** Campos não inicializados têm valor zero (0 para números, "" para strings, `nil` para ponteiros).

```go
var tree Tree
// tree.ID = 0
// tree.Species = ""
// tree.Height = 0.0
// tree.Diameter = 0.0
```

**Formas de criar structs:**
```go
// 1. Literal completo
tree1 := Tree{
    ID:       1,
    Species:  "Eucalyptus",
    Height:   12.0,
    Diameter: 18.5,
}

// 2. Literal parcial (resto fica zero value)
tree2 := Tree{
    Species: "Quercus",
}

// 3. Zero value
var tree3 Tree

// 4. Com new (retorna ponteiro)
tree4 := new(Tree)  // *Tree (ponteiro)
tree4.Species = "Castanea"
```

### Ponteiros para Structs
```go
// Criar ponteiro
tree := &Tree{
    Species: "Pinus",
    Height:  10.0,
}
// tree é *Tree (ponteiro)

// Acesso automático
tree.Height = 15.0  // Go faz (*tree).Height automaticamente!

tree := &Tree{Height: 10.0}

// Estas são EQUIVALENTES:
tree.Height = 20.0      // ✅ Go simplifica
(*tree).Height = 20.0   // ✅ Forma explícita (raramente usada)
```

### Methods
* Methods são funções associadas a um tipo.
* Podem ter receiver por valor ou por ponteiro. Isto é, se o receiver é por valor, o método recebe uma cópia do struct. Se o receiver é por ponteiro, o método recebe um ponteiro para o struct, permitindo modificar o original.

```go
type Tree struct {
    Height   float64
    Diameter float64
}

// Method com VALUE receiver (cópia)
func (t Tree) Volume() float64 {
    // Fórmula simplificada: cilindro
    radius := t.Diameter / 2.0 / 100.0  // cm para metros
    return 3.14 * radius * radius * t.Height
}

// Method com POINTER receiver (referência)
func (t *Tree) Grow(amount float64) {
    t.Height += amount  // modifica o original
}

// Uso
tree := Tree{Height: 10.0, Diameter: 20.0}
fmt.Println(tree.Volume())  // Calcula volume (cópia)

tree.Grow(2.0)  // Modifica altura (original)
fmt.Println(tree.Height)  // 12.0

```

### Value vs Pointer Receivers
* **Consistência é importante:** se um tipo tem um método com pointer receiver, é recomendado que todos os métodos usem pointer receivers para evitar confusão.
* **Regra de ouro:** Se UM method precisa de pointer receiver, TODOS devem usar pointer receiver (consistência).

**Value Receiver `(t Tree)`:**
```go
func (t Tree) GetHeight() float64 {
    return t.Height  // retorna altura (cópia), só lê
}
```

Usa quando:
* ✅ Method só lê dados (não modifica)
* ✅ Struct é pequena (poucos campos)
* ✅ Struct tem apenas tipos primitivos

**Pointer Receiver `(t *Tree)`:**
```go
func (t *Tree) SetHeight(h float64) {
    t.Height = h  // modifica altura (original)
}
```

Usa quando:
* ✅ Method modifica a struct
* ✅ Struct é grande (muitos campos, evita cópia)
* ✅ Struct tem campos que não devem ser copiados (mutexes, etc)

Exemplo comparativo:
```go
type Tree struct {
    Height   float64
}

// ❌ Value receiver - NÃO modifica
func (t Tree) GroWrong(meters float64) {
    t.Height += meters  // modifica cópia, não o original
}

// ✅ Pointer receiver - modifica
func (t *Tree) GrowRight(meters float64) {
    t.Height += meters  // modifica o original
}

func main() {
    tree := Tree{Height: 10.0}

    tree.GroWrong(5.0)
    fmt.Println(tree.Height)  // 10.0 (não mudou)

    tree.GrowRight(5.0)
    fmt.Println(tree.Height)  // 15.0 (mudou!)
}
```

### Convenções

```go
// Models (structs de dados)
type Tree struct {
    ID        int       `json:"id"`
    Species   string    `json:"species"`
    Height    float64   `json:"height"`
    Diameter  float64   `json:"diameter"`
    CreatedAt time.Time `json:"created_at"`
}

// Tags JSON - importante para marshal/unmarshal
```

**Methods típicos:**
```go
// Validação
func (t *Tree) Validate() error {
    if t.Height < 0 {
        return errors.New("altura não pode ser negativa")
    }
    if t.Diameter < 0 {
        return errors.New("diâmetro não pode ser negativo")
    }
    return nil
}

// Cálculos
func (t Tree) Volume() float64 {
    // tua fórmula aqui
}

func (t Tree) BasalArea() float64 {
    // área basal
}
```