package funcs

import "fmt"

type Pessoas struct {
	Nome  string
	Idade int
}

func pessoa(pessoa Pessoas, name string, age int) Pessoas {
	pessoa.Nome = name
	pessoa.Idade = age

	fmt.Printf("Nome: %s, Idade: %d\n", pessoa.Nome, pessoa.Idade)
	return pessoa
}

func Funcs() {
	fmt.Println("=== Funções em Go ===")

	p := Pessoas{}
	p = pessoa(p, "João", 25)
	p = pessoa(p, "Maria", 30)


	fmt.Println("Struct final:", p)
}
