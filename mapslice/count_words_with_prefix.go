package mapslice

// CountWordsWithPrefix возвращает количество строк, которые начинаются с prefix. Пустой prefix подходит ко всем строкам.
func CountWordsWithPrefix(words []string, prefix string) int {
	// TODO: реализовать функцию.
	result := 0
	if prefix == "" {
		return len(words)
	}
	for _, word := range words {
		if len(word) < len(prefix) {
			continue
		}

		if word[:len(prefix)] == prefix {
			result++
		}
	}
	return result
}
