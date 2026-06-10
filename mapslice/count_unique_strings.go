package mapslice

// CountUniqueStrings возвращает количество уникальных строк в words.
func CountUniqueStrings(words []string) int {
	// TODO: реализовать функцию.
	seen := make(map[string]int)
	result := 0
	for _, value := range words {
		seen[value]++
	}
	result = len(seen)

	return result
}
