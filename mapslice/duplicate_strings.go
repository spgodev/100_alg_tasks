package mapslice

// DuplicateStrings возвращает строки, которые встречаются больше одного раза, в порядке их первого повторного появления.
func DuplicateStrings(words []string) []string {
	// TODO: реализовать функцию.
	result := make([]string, 0)
	seen := make(map[string]int)

	for _, word := range words {
		seen[word]++
		if seen[word] == 2 {
			result = append(result, word)
		}
	}
	return result
}
