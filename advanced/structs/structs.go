package structs

import (
	"fmt"
)

func Structs () {
	fmt.Println("=== Structs em Go ===")

	type Raca struct {
		Nome string
		Origem string
		Caracteristicas []string
	}
	type Cachorro struct {
		Nome string
		Idade int
		Raca Raca
	}

	// Criando uma instância
	cachorro1 := Cachorro{
		Nome: "Rex",
		Idade: 5,
		Raca: Raca{
			Nome: "Pastor Alemão",
			Origem: "Alemanha",
			Caracteristicas: []string{"Inteligente", "Leal", "Protetor"},
		},
	}

	// Acessando os campos
	fmt.Printf("Nome: %s\n", cachorro1.Nome)
	fmt.Printf("Idade: %d\n", cachorro1.Idade)
	fmt.Printf("Raça: %s\n", cachorro1.Raca.Nome)
	fmt.Printf("Origem: %s\n", cachorro1.Raca.Origem)
	fmt.Printf("Características: %v\n", cachorro1.Raca.Caracteristicas)
}

