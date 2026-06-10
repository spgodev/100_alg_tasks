package mapslice

// KeepAllowedInts возвращает элементы src, для которых allowed[value] == true.
func KeepAllowedInts(src []int, allowed map[int]bool) []int {
	// TODO: реализовать функцию.
	result := make([]int, 0)
	for _, value := range src {
		if allowed[value] {
			result = append(result, value)
		}
	}
	return result
}
