package mapslice

// CountWordsByLastLetter считает непустые строки по последнему байту. Пустые строки учитываются по ключу "".
func CountWordsByLastLetter(words []string) map[string]int {
	// TODO: реализовать функцию.
	result := make(map[string]int)
	for _, word := range words {
		if word == "" {
			result[""]++
			continue
		}
		lastLetter := string(word[len(word)-1])
		result[lastLetter]++
	}
	return result
}
