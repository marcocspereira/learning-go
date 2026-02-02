// Sistema de Validação de Passwords
//
// Cria funções para validar passwords com as seguintes regras:
// Funções a implementar:
//
// validaComprimento(password string) (bool, string)
// - Retorna true se >= 8 caracteres
// - String de erro explicativa se falhar
//
// contemNumero(password string) bool
// - Retorna true se contém pelo menos 1 dígito (0-9)
//
// contemMaiuscula(password string) bool
// - Retorna true se contém pelo menos 1 maiúscula
//
// validaPassword(password string) (bool, []string)
// - Usa as 3 funções acima
// - Retorna true + slice vazio se válida
// - Retorna false + slice com todos os erros se inválida
//
// Requisitos técnicos:
// - Usa múltiplos returns
// - Para verificar caracteres, usa um for com range sobre a string
// - Pesquisa: unicode.IsDigit() e unicode.IsUpper() do package unicode

// Exemplo de uso:
/*
func main() {
    valida, erros := validaPassword("abc")
    if !valida {
        fmt.Println("Password inválida:")
        for _, erro := range erros {
            fmt.Println("-", erro)
        }
    } else {
        fmt.Println("Password válida!")
    }
}
```

Output esperado para "abc":
```
Password inválida:
- Password deve ter no mínimo 8 caracteres
- Password deve conter pelo menos um número
- Password deve conter pelo menos uma letra maiúscula
```

Output esperado para "SecurePass123":
```
Password válida!
*/

package main

import (
	"fmt"
	"unicode"
)

/*
*
Utiliza *string para retornar mensagens de erro opcionais como nil quando não há erro. Utiliza um tipo de dado ponteiro para distinguir entre ausência de erro e string vazia para efeitos académicos.
A alternativa seria retornar apenas string e usar string vazia para indicar ausência de erro.
*/
func validaComprimento(password string) (bool, *string) {
	/*
	   len() counts bytes, not runes (characters)
	   []rune converts the string to a slice of runes and a rune represents a Unicode code point
	*/
	if len([]rune(password)) < 8 {
		errMsg := "Password deve ter no mínimo 8 caracteres"
		return false, &errMsg
	}
	return true, nil
}

func contemNumero(password string) bool {
	for _, char := range []rune(password) {
		// IsDigit reports whether the rune is a decimal digit.
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}

func contemMaiuscula(password string) bool {
	for _, char := range []rune(password) {
		// IsUpper reports whether the rune is an upper case letter.
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}

func validaPassword(password string) (bool, []string) {
	var erros []string

	if valid, errMsg := validaComprimento(password); !valid {
		erros = append(erros, *errMsg)
	}
	if !contemNumero(password) {
		erros = append(erros, "Password deve conter pelo menos um número")
	}
	if !contemMaiuscula(password) {
		erros = append(erros, "Password deve conter pelo menos uma letra maiúscula")
	}

	return len(erros) == 0, erros

	return false, erros
}

func main() {
	// Exemplo de uso
	valida, erros := validaPassword("abc")
	if !valida {
		fmt.Println("Password inválida:")
		for _, erro := range erros {
			fmt.Println("-", erro)
		}
	} else {
		fmt.Println("Password válida!")
	}
}
