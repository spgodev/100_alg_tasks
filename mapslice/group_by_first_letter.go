package mapslice

// GroupByFirstLetter группирует непустые строки по первому байту. Пустые строки попадают в ключ "".
func GroupByFirstLetter(words []string) map[string][]string {
	// TODO: реализовать функцию.
	result := make(map[string][]string)
	for _, word := range words {
		if word == "" {
			result[""] = append(result[""], "")
			continue
		}
		firstLetter := word[:1]
		result[firstLetter] = append(result[firstLetter], word)
	}

	return result
}
