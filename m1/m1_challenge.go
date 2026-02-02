/**
Cria um programa que analisa logs de um servidor web e gera estatísticas.

Input (slice de strings):
logs := []string{
    "2024-01-15 10:23:45 GET /home 200",
    "2024-01-15 10:24:12 POST /api/users 201",
    "2024-01-15 10:25:33 GET /products 200",
    "2024-01-15 10:26:01 GET /home 304",
    "2024-01-15 10:27:15 DELETE /api/users/5 404",
    "2024-01-15 10:28:42 GET /about 500",
    "2024-01-15 10:29:18 POST /api/login 200",
}

### Requisitos:

**1. Função `parseLinha(linha string) (metodo, caminho string, codigo int, erro error)`**
- Separa os componentes de cada linha
- Retorna erro se formato inválido
- Dica: pesquisa `strings.Fields()` ou `strings.Split()`

**2. Função `contarPorMetodo(logs []string) map[string]int`**
- Retorna mapa: `{"GET": 4, "POST": 2, "DELETE": 1}`
- Ignora linhas inválidas

**3. Função `contarErros(logs []string) (total int, porcentagem float64)`**
- Conta códigos >= 400
- Calcula percentagem de erros

**4. Função `caminhoMaisAcedido(logs []string) string`**
- Retorna o caminho com mais acessos
- Se empate, retorna qualquer um

**5. `main()` que imprime**:
```
=== Estatísticas de Logs ===
Total de requisições: 7

Requisições por método:
- GET: 4
- POST: 2
- DELETE: 1

Erros: 2 (28.57%)

Caminho mais acedido: /home (2 acessos)

Conceitos a aplicar:

✅ Múltiplos returns
✅ Error handling
✅ Loops com range
✅ Slices
✅ Maps (novo! pesquisa como usar)
✅ Funções modulares
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseLinha(linha string) (metodo, caminho string, codigo int, erro error) {
	fields := strings.Fields(linha)
	if len(fields) != 5 {
		erro = fmt.Errorf("formato inválido, precisa de 5 campos e foram encontrados %d", len(fields))
		return // retorna zero values e o erro: metodo="", caminho="", codigo=0, erro="formato inválido"
	}
	metodo = fields[2]
	caminho = fields[3]
	codigo, erro = strconv.Atoi(fields[4])
	if erro != nil {
		erro = fmt.Errorf("código de status HTTP inválido %s: %w", fields[4], erro)
	}
	return
}

func contarErros(logs []string) (total int, porcentagem float64) {
	totalRequisicoesValidas := 0
	for _, linha := range logs {
		_, _, codigo, erro := parseLinha(linha)
		if erro != nil {
			continue // ignora linhas inválidas
		}
		totalRequisicoesValidas++

		if codigo >= 400 {
			total++
		}
	}

	if totalRequisicoesValidas == 0 {
		return 0, 0.0
	}

	porcentagem = (float64(total) / float64(totalRequisicoesValidas)) * 100
	return
}

func contarPorMetodo(logs []string) map[string]int {
	contagem := make(map[string]int)
	for _, linha := range logs {
		metodo, _, _, erro := parseLinha(linha)
		if erro != nil {
			continue // ignora linhas inválidas
		}
		contagem[metodo]++
	}
	return contagem
}

func caminhoMaisAcedido(logs []string) string {
	contagem := make(map[string]int)
	for _, linha := range logs {
		_, caminho, _, erro := parseLinha(linha)
		if erro != nil {
			continue // ignora linhas inválidas
		}
		contagem[caminho]++
	}
	maxAcessos := 0
	misAcedido := ""
	for caminho, acessos := range contagem {
		if acessos > maxAcessos {
			maxAcessos = acessos
			misAcedido = caminho
		}
	}
	return misAcedido
}

func main() {

	logs := []string{
		"2024-01-15 10:23:45 GET /home 200",
		"2024-01-15 10:24:12 POST /api/users 201",
		"2024-01-15 10:25:33 GET /products 200",
		"2024-01-15 10:26:01 GET /home 304",
		"2024-01-15 10:27:15 DELETE /api/users/5 404",
		"2024-01-15 10:28:42 GET /about 500",
		"2024-01-15 10:29:18 POST /api/login 200",
	}
	totalRequisicoes := len(logs)
	contagemMetodos := contarPorMetodo(logs)
	totalErros, porcentagemErros := contarErros(logs)
	caminho := caminhoMaisAcedido(logs)

	fmt.Println("=== Estatísticas de Logs ===")
	fmt.Printf("Total de requisições: %d\n\n", totalRequisicoes)

	fmt.Println("Requisições por método:")
	for metodo, contagem := range contagemMetodos {
		fmt.Printf("- %s: %d\n", metodo, contagem)
	}

	fmt.Printf("\nErros: %d (%.2f%%)\n\n", totalErros, porcentagemErros)
	fmt.Printf("Caminho mais acedido: %s\n", caminho)
}
