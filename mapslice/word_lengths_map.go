package mapslice

// WordLengthsMap возвращает map, где ключ — слово, значение — длина слова в байтах. При повторах слово остаётся одним ключом.
func WordLengthsMap(words []string) map[string]int {
	// TODO: реализовать функцию.
	result := make(map[string]int)
	for i, v := range words {
		result[v] = len(words[i])
	}
	return result
}
