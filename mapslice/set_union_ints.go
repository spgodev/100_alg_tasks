package mapslice

// SetUnionInts возвращает объединение уникальных значений: сначала новые элементы из a, затем новые элементы из b.
func SetUnionInts(a []int, b []int) []int {
	// TODO: реализовать функцию.
	added := make(map[int]bool)
	result := make([]int, 0)

	for _, value := range a {
		if !added[value] {
			result = append(result, value)
			added[value] = true
		}
	}

	for _, value := range b {
		if !added[value] {
			result = append(result, value)
			added[value] = true
		}
	}

	return result
}
