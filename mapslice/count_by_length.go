package mapslice

// CountByLength считает, сколько строк каждой длины встречается в words.
func CountByLength(words []string) map[int]int {
	// TODO: реализовать функцию.
	result := make(map[int]int)
	for _, word := range words {
		length := len(word)
		result[length]++

	}
	return result
}
