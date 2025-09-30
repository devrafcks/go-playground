package main

import (
	"fmt"
	"controles/if-else"
	"controles/forLace"
)

func main() {
	num := 11
	fmt.Println("=== Verificação de Número ===")
	fmt.Printf("Número: %d\n", num)
	ifelse.CheckNumber(num)

	idade := 20
	fmt.Println("\n=== Verificação de Idade ===")
	fmt.Printf("Idade: %d\n", idade)
	ifelse.CheckIdade(idade)

	nota := 6.5
	fmt.Println("\n=== Verificação de Nota ===")
	fmt.Printf("Nota: %.2f\n", nota)
	ifelse.CheckNota(nota)

	numPrimo := 7
	fmt.Println("\n=== Verificação de Número Primo ===")
	fmt.Printf("Número: %d\n", numPrimo)
	ifelse.CheckNumeroPrimo(numPrimo)

	fmt.Println("\n=== Impressão de Números ===")
	forLace.PrintNumbers()
}
