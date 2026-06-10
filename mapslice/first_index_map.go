package mapslice

// FirstIndexMap возвращает map значение -> индекс первого появления этого значения.
func FirstIndexMap(src []int) map[int]int {
	// TODO: реализовать функцию.
	result := make(map[int]int)
	for i, value := range src {
		if _, ok := result[value]; !ok {
			result[value] = i
		}
	}
	return result
}
