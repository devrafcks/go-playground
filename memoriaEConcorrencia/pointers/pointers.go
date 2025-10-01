package pointers

import "fmt"

func Pointer() {
	a := 10
	b := &a 
	fmt.Println("Valor de a:", a)
	fmt.Println("Endereço de a:", &a)
	fmt.Println("Valor de b:", b)
	fmt.Println("Valor apontado por b:", *b)
}