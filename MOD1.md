# Módulo 1

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

