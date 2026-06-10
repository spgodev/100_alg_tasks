package mapslice

// PositionsMap возвращает map значение -> все индексы этого значения в порядке возрастания.
func PositionsMap(src []int) map[int][]int {
	// TODO: реализовать функцию.
	result := make(map[int][]int)
	for i, v := range src {
		result[v] = append(result[v], i)
	}
	return result
}
