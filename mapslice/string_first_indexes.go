package mapslice

// StringFirstIndexes возвращает map строка -> индекс первого появления.
func StringFirstIndexes(words []string) map[string]int {
	// TODO: реализовать функцию.
	result := make(map[string]int)
	for i, w := range words {
		if _, ok := result[w]; !ok {
			result[w] = i
		}
	}
	return result
}
