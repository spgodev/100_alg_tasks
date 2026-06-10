package mapslice

// RemoveStopWords возвращает words без строк, для которых stop[word] == true. Порядок сохраняется.
func RemoveStopWords(words []string, stop map[string]bool) []string {
	// TODO: реализовать функцию.
	result := make([]string, 0, len(words))
	for _, word := range words {
		if !stop[word] {
			result = append(result, word)
		}
	}
	return result
}
