package mapslice

// GroupByParity группирует числа по чётности в ключи "even" и "odd". Порядок внутри групп сохраняется.
func GroupByParity(src []int) map[string][]int {
	// TODO: реализовать функцию.
	result := map[string][]int{
		"even": {},
		"odd":  {},
	}

	for _, value := range src {
		if value%2 == 0 {
			result["even"] = append(result["even"], value)
		}
		if value%2 == 1 || value%2 == -1 {
			result["odd"] = append(result["odd"], value)
		}
	}
	return result
}
