package mapslice

// MapFromPairs создаёт map из параллельных слайсов keys и values. Если длины разные, используется меньшая длина.
func MapFromPairs(keys []string, values []int) map[string]int {
	// TODO: реализовать функцию.
	result := make(map[string]int)
	limit := len(keys)

	if len(values) < limit {
		limit = len(values)
	}

	for i := 0; i < limit; i++ {
		result[keys[i]] = values[i]
	}
	return result
}
