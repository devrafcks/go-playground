package ifelse

import "fmt"

func CheckNumber(num int) {
	if num%2 == 0 {
		fmt.Println(num, "par")
	} else {
		fmt.Println(num, "ímpar")
	}
}

func CheckIdade(idade int) {
	if idade < 18 {
		fmt.Println("Menor de idade")
	} else {
		fmt.Println("Maior de idade")
	}
}

func CheckNota(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado")
	}
	if nota < 7 && nota >= 5 {
		fmt.Println("Recuperação")
	} else {
		fmt.Println("Reprovado")
	}
}

func CheckNumeroPrimo(num int) {
	if num <= 1 {
		fmt.Println(num, "não é primo")
		return
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			fmt.Println(num, "não é primo")
			return
		}
	}
	fmt.Println(num, "é primo")
}