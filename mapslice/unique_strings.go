package mapslice

// UniqueStrings возвращает уникальные строки в порядке первого появления.
func UniqueStrings(words []string) []string {
	// TODO: реализовать функцию.
	seen := make(map[string]bool)
	result := make([]string, 0)
	for _, w := range words {
		if !seen[w] {
			result = append(result, w)
			seen[w] = true
		}
	}
	return result
}
