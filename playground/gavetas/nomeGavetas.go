package gavetas

func ObjetoDaGaveta() string {
	var objeto = map[string]string{
		"nome":  "caneta",
		"cor":   "azul",
		"tinta": "preta",
	}
	return objeto["nome"] + ", " + objeto["cor"] + ", " + objeto["tinta"]
}

func NomeDaGaveta() string {
	var nome = map[string]string{
		"escritorio": "Gaveta de Escritório",
		"cozinha":    "Gaveta de Cozinha",
		"banheiro":   "Gaveta de Banheiro",
	}

	return nome["escritorio"] + ", " + nome["cozinha"] + ", " + nome["banheiro"]
}