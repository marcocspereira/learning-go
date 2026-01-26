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
