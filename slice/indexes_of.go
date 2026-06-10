package slice

// IndexesOf возвращает все индексы, на которых встречается target.
func IndexesOf(src []int, target int) []int {
	// TODO: реализовать функцию.
	result := make([]int, 0, len(src))
	for i := range src {
		if src[i] == target {
			result = append(result, i)
		}
	}
	return result
}
