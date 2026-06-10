package mapslice

// IndexWordsByLength возвращает map длина слова -> индексы слов этой длины.
func IndexWordsByLength(words []string) map[int][]int {
	// TODO: реализовать функцию.
	result := make(map[int][]int)
	for i, word := range words {
		result[(len(word))] = append(result[(len(word))], i)
	}
	return result
}
