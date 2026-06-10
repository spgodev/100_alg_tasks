package mapslice

// LastIndexMap возвращает map значение -> индекс последнего появления этого значения.
func LastIndexMap(src []int) map[int]int {
	// TODO: реализовать функцию.
	result := make(map[int]int)

	for i, value := range src {
		result[value] = i
	}

	return result
}
