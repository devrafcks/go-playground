package rangeStructure

import "fmt"

func RangeStructure (){
	// Range em Slice
	fmt.Println("=== Slice ===")
	numeros := []int{10, 20, 30}
	for i, v := range numeros {
		fmt.Printf("Índice: %d, Valor: %d\n", i, v)
	}

	// Ignorando o índice
	for _, v := range numeros {
		fmt.Printf("Valor sem índice: %d\n", v)
	}

	// Range em Array
	fmt.Println("\n=== Array ===")
	letras := [3]string{"A", "B", "C"}
	for i, l := range letras {
		fmt.Printf("Posição: %d, Letra: %s\n", i, l)
	}

	// Range em String
	fmt.Println("\n=== String ===")
	palavra := "GoLang"
	for i, c := range palavra {
		fmt.Printf("Posição: %d, Rune: %d, Caractere: %c\n", i, c, c)
	}

	// Range em Map
	fmt.Println("\n=== Map ===")
	paises := map[string]string{
		"BR": "Brasil",
		"US": "Estados Unidos",
		"FR": "França",
	}
	for chave, valor := range paises {
		fmt.Printf("Chave: %s, Valor: %s\n", chave, valor)
	}

	// Range em Channel
	fmt.Println("\n=== Channel ===")
	ch := make(chan int, 3)
	ch <- 100
	ch <- 200
	ch <- 300
	close(ch) 

	for v := range ch {
		fmt.Printf("Recebido: %d\n", v)
	}
}