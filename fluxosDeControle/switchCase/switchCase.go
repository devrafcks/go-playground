package switchcase

import "fmt"

func PrintDayOfWeek(day int) {
	switch day {
	case 1:
		fmt.Println("Domingo")
	case 2:
		fmt.Println("Segunda-feira")
	case 3:
		fmt.Println("Terça-feira")
	case 4:
		fmt.Println("Quarta-feira")
	case 5:
		fmt.Println("Quinta-feira")
	case 6:
		fmt.Println("Sexta-feira")
	case 7:
		fmt.Println("Sábado")
	default:
		fmt.Println("Dia inválido")
	}
}

func PrintFruitColor(fruit string) {
	switch fruit {
	case "maçã":
		fmt.Println("Vermelha")
	case "banana":
		fmt.Println("Amarela")
	case "laranja":
		fmt.Println("Laranja")	
	case "uva":
		fmt.Println("Roxa")
	default:
		fmt.Println("Cor desconhecida")
	}	
}

func PrintMonthName(month int) {
	switch month {
	case 1:
		fmt.Println("Janeiro")
	case 2:
		fmt.Println("Fevereiro")
	case 3:
		fmt.Println("Março")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Maio")
	case 6:
		fmt.Println("Junho")
	case 7:
		fmt.Println("Julho")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Setembro")
	case 10:
		fmt.Println("Outubro")
	case 11:
		fmt.Println("Novembro")
	case 12:
		fmt.Println("Dezembro")
	default:
		fmt.Println("Mês inválido")
	}
}
