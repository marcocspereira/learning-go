// Analisador de Temperaturas
// Cria um programa que:

// Recebe um slice de temperaturas (em Celsius) da semana
// Calcula e imprime:

// Temperatura média
// Temperatura máxima e em que dia ocorreu
// Temperatura mínima e em que dia ocorreu
// Quantos dias tiveram temperatura acima de 25°C

// Requisitos:
// - Usa `for` com `range`
// - Usa `len()` para calcular média
// - **Não uses bibliotecas externas** (faz tudo manualmente)

// Output esperado:
// Temperatura média: 24.36°C
// Máxima: 28.00°C (dia 5)
// Mínima: 21.00°C (dia 7)
// Dias acima de 25°C: 3

package main

import "fmt"

func main() {
	temperaturas := []float64{22.5, 24.0, 26.5, 23.0, 28.0, 25.5, 21.0}
	maxTemp := temperaturas[0]
	maxDay := 1
	minTemp := temperaturas[0]
	minDay := 1
	daysAbove25 := 0
	averageTemp := 0.0

	for i, temp := range temperaturas {
		averageTemp += temp
		if temp > maxTemp {
			maxTemp = temp
			maxDay = i + 1
		}
		if temp < minTemp {
			minTemp = temp
			minDay = i + 1
		}
		if temp > 25.0 {
			daysAbove25++
		}
	}

	averageTemp /= float64(len(temperaturas))

	fmt.Printf("Temperatura média: %.2f°C\n", averageTemp)
	fmt.Printf("Máxima: %.2f°C (dia %d)\n", maxTemp, maxDay)
	fmt.Printf("Mínima: %.2f°C (dia %d)\n", minTemp, minDay)
	fmt.Printf("Dias acima de 25°C: %d\n", daysAbove25)
}
