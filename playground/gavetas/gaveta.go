package gavetas

func QuantidadeDeGavetas(x int) int {
	var gavetas []int
	for i := 0; i < x; i++ {
		gavetas = append(gavetas, i)
	}
	return len(gavetas)
}