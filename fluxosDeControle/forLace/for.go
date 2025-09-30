package forLace

import "fmt"

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Número:", i)
	}

	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Println("Soma dos números de 1 a 5 é:", sum)

}
