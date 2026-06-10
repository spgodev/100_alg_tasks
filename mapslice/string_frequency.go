package mapslice

// StringFrequency возвращает map строка -> количество вхождений.
func StringFrequency(words []string) map[string]int {
	// TODO: реализовать функцию.
	result := make(map[string]int)
	for _, w := range words {
		result[w]++
	}
	return result
}
