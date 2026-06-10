package mapslice

// SymmetricDifferenceInts возвращает уникальные значения, которые встречаются только в одном из двух слайсов. Сначала порядок из a, затем из b.
func SymmetricDifferenceInts(a []int, b []int) []int {
	// TODO: реализовать функцию.
	inA := make(map[int]bool)
	inB := make(map[int]bool)
	added := make(map[int]bool)
	result := make([]int, 0)
	for _, value := range a {
		inA[value] = true
	}

	for _, value := range b {
		inB[value] = true
	}
	for _, value := range a {
		if !inB[value] && !added[value] {
			result = append(result, value)
			added[value] = true
		}
	}
	for _, value := range b {
		if !inA[value] && !added[value] {
			result = append(result, value)
			added[value] = true
		}
	}
	return result
}
