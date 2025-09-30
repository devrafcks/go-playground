package main

import (
	"fmt"
	"hello/meet"
	"hello/operations"
	"hello/gavetas"
)

func main() {
	x := meet.GetRandomNumber()
	meet.PrintNumber(x)
	fmt.Println("sum of 1 and 2 is:", operations.Add(1, 2))
	fmt.Println("difference of 5 and 3 is:", operations.Subtract(5, 3))
	fmt.Println("product of 2 and 3 is:", operations.Multiply(2, 3))
	fmt.Println("division of 10 by 2 is:", operations.Divide(10, 2))
	fmt.Println("Quantidade de Gavetas é:", gavetas.QuantidadeDeGavetas(12))
	fmt.Println("Objetos na Gaveta é:", gavetas.ObjetoDaGaveta())
	fmt.Println("Nomes das Gavetas são:", gavetas.NomeDaGaveta())
	fmt.Println("End of program")
}