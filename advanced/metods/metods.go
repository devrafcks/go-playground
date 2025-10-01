package metods

import "fmt"

type Pessoas struct {
	Nome  string
	Idade int
}

func (p *Pessoas) SetNome(nome string) {
	p.Nome = nome
}

func (p *Pessoas) SetIdade(idade int) {
	p.Idade = idade
}

func (p Pessoas) GetNome() string {
	return p.Nome
}

func (p Pessoas) GetIdade() int {
	return p.Idade
}

func Metods() {
	fmt.Println("=== Métodos em Go ===")

	p := Pessoas{}
	p.SetNome("João")
	p.SetIdade(25)

	fmt.Println("Nome:", p.GetNome())
	fmt.Println("Idade:", p.GetIdade())
}
