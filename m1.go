// Calculadora de IMC
//
// Cria um programa que:
//
// Declara variáveis para peso (kg) e altura (metros)
// Calcula o IMC (peso / altura²)
// Classifica o resultado:
//
// < 18.5: "Abaixo do peso"
// 18.5-24.9: "Peso normal"
// 25-29.9: "Sobrepeso"
//
// = 30: "Obesidade"
//
// Imprime o IMC com 2 casas decimais e a classificação

package main

import "fmt"

// go run m1.go
func main() {
	peso := 75.0
	altura := 1.71
	imc := peso / (altura * altura)

	var classificacao string

	if imc < 18.5 {
		classificacao = "Abaixo do peso"
	} else if imc < 25 {
		classificacao = "Peso normal"
	} else if imc < 30 {
		classificacao = "Sobrepeso"
	} else {
		classificacao = "Obesidade"
	}

	fmt.Printf("IMC: %.2f - %s\n", imc, classificacao)
}
