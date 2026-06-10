package mapslice

// SameStringFrequencies возвращает true, если в a и b одинаковые частоты строк.
func SameStringFrequencies(a []string, b []string) bool {
	// TODO: реализовать функцию.
	if len(a) != len(b) {
		return false
	}

	counts := make(map[string]int)

	for _, word := range a {
		counts[word]++
	}

	for _, word := range b {
		if counts[word] == 0 {
			return false
		}
		counts[word]--
	}

	return true
}
