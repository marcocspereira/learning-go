# Módulo 1

## Módulo 1.1: Tipos Básicos e Declarações

* Go é estaticamente tipado: as variáveis devem ter seus tipos definidos em tempo de compilação.
* `:=` é usado para declarar e inicializar variáveis de forma concisa, logo mais idimática.
* `var` é usado para declarações explícitas de variáveis, especialmente fora de funções.
* Variáveis não utilizadas geram erros de compilação. Quando usar o `var` dentro de funções:
  * Quando queres zero value explícito (`var count int`, `0`, mais claro que count `:= 0`)
  * Quando o tipo não é óbvio pela atribuição (`var timeout time.Duration = 30 * time.Second`)
* **Zero values** são atribuídos automaticamente a variáveis não inicializadas (`0` para inteiros, `""` para strings, `false` para booleanos, `nil` para ponteiros, slices, maps, channels e interfaces).
  * `var classificacao string  // "" (string vazia)`
  * `var count int            // 0`
  * `var active bool          // false`
  * `var ptr *int             // nil`

### Tipos básicos

* `int`, `float64`, `string`, `bool`
* Conversão de tipos é explícita: `float64(x)` converte `x` para `float64`.
* Strings são imutáveis e suportam concatenação com `+`.

Diferença entre int8, int16, int32, int64:
* `int8`: -128 a 127 (`byte` é um alias para `int8`)
* `int16`: -32,768 a 32,767
* `int32`: -2,147,483,648 a 2,147,483,647 (`rune` é um alias para `int32`, representa um ponto de código Unicode)
* `int64`: -9,223,372,036,854,775,808 a 9,223,372,036,854,775,807

Diferença entre float32 e float64:
* `float32`: precisão simples, 7 dígitos decimais de precisão. Só é usado quando há necessidade de economizar memória. (array grandes, por exemplo)
* `float64`: precisão dupla, 15-16 dígitos decimais de precisão. Literal padrão para números de ponto flutuante em Go.

```go
// Módulo 1

// Declaração explícita com tipo
var name string = "Anabela"
var age int = 38

// Declaração com inferência de tipo
var city = "Porto"

// Short declaration (apenas dentro de funções)
// country := "Portugal"

// constante
const pi = 3.14
```

---

## Módulo 1.2: Loops e Slices

### Loops
Em Go, o loop `for` é a única estrutura de repetição. Pode ser usado de várias formas:
```go
// Loop tradicional
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// Loop estilo while (sem inicialização e pós-execução)
for i < 10 {
    fmt.Println(i)
    i++
}

// Loop infinito
for {
    // fazer algo
    // break para sair do loop
}

// Iterar sobre slices/arrays
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Println("Index:", index, "Value:", value)
}

// Iterar sobre strings (caracteres Unicode)
str := "Olá"
for index, char := range str {
    fmt.Printf("Index: %d, Char: %c\n", index, char)
}

// Ignorar o índice ou valor
for _, value := range numbers {
    fmt.Println("Value:", value)
}
```

### Slices vs Arrays

```go
var arr [5]int = [5]int{1, 2, 3, 4, 5} // Array com tamanho fixo

slice := []int{1, 2, 3, 4, 5} // Slice dinâmico
slice = append(slice, 6) // Adiciona elemento ao slice
lenSlice := len(slice) // Obtém o tamanho do slice
capSlice := cap(slice) // Obtém a capacidade do slice
numSlice := slice[1:3] // Sub-slice (elementos do índice 1 ao 2, ou seja os valores 2 e 3)
```

* 99% dos casos, usar slices é preferível devido à sua flexibilidade.
* Slices são passados por referência, ou seja, alterações numa função afetam o slice original. Exemplo:
```go
func modifySlice(s []int) {
    s[0] = 100
}
mySlice := []int{1, 2, 3}
modifySlice(mySlice)
fmt.Println(mySlice) // Output: [100, 2, 3]
```
* Para iniciar um slice com tamanho e capacidade específicas, usar `make`:
```go
// make(tipo, length, capacity)
s1 := make([]string, 0, 10)   // len=0, cap=10 (vazio mas com espaço reservado)
s2 := make([]int, 5)          // len=5, cap=5 (inicializado com zero values)
s3 := make([]string, 5, 10) // len=5, cap=10 (inicializado com zero values)
```
* Diferença entre length e capacity:
  * `len(slice)`: número de elementos atualmente no slice.
  * `cap(slice)`: número máximo de elementos que o slice pode conter antes de precisar realocar memória.

* Quando usar slice literal vs make:
  * Usar slice literal (`[]int{1,2,3}`) quando se conhece os valores iniciais.
  * Usar `make` quando se quer criar um slice vazio com capacidade pré-definida ou quando se quer inicializar com zero values.
  ```go
  // Slice literal
  numbers := []int{1, 2, 3, 4, 5}
  // // make - quando vais popular depois
  results := make([]string, 0, 100)  // reserva espaço para evitar realocações
  for i := 0; i < 100; i++ {
    results = append(results, processData(i))
  }
  ```

* **Use cases para arrays:**
  * Quando o tamanho é fixo e conhecido em tempo de compilação.
  * Para otimizações de desempenho em casos específicos (evitar overhead de slices).
* **Use cases para slices:**
  * Quando o tamanho pode variar.
  * Para manipulação dinâmica de coleções de dados.


---

## Módulo 1.3: Funções e Múltiplos Retornos

### Características das Funções em Go
```go
// Declaração de função simples
func add(a int, b int) int {
    return a + b
}
sum := add(3, 5) // sum = 8

// Função com parâmetros do mesmo tipo (sintaxe concisa)
func multiply(a, b int) int {
    return a * b
}
mul := multiply(4, 6) // mul = 24

// Função com múltiplos retornos
func divideAndRemainder(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

// Chamada de função com múltiplos retornos
q, r := divideAndRemainder(10, 3)
fmt.Println("Quociente:", q, "Resto:", r) // Output: Quociente: 3 Resto: 1

// Retorno nomeado
func swap(x, y string) (first string, second string) {
    first = y
    second = x
    return
}
s1, s2 := swap("hello", "world")
fmt.Println(s1, s2) // Output: world hello

// Função variádica (aceita número variável de argumentos)
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}
totalSum := sum(1, 2, 3, 4, 5) // totalSum = 15
```

### Error Handling
Em Go, o tratamento de erros é feito retornando um valor de erro como o último retorno da função. Não há exceções como em outras linguagens.

```go
import "fmt"
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("divisão por zero não é permitida")
    }
    return a / b, nil
}

result, err := divide(10, 0)
if err != nil {
    fmt.Println("Erro:", err)
    return
}
fmt.Println("Resultado:", result)
```