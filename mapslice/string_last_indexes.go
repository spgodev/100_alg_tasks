package mapslice

// StringLastIndexes возвращает map строка -> индекс последнего появления.
func StringLastIndexes(words []string) map[string]int {
	// TODO: реализовать функцию.
	result := make(map[string]int)
	for i, w := range words {
		result[w] = i
	}
	return result
}
