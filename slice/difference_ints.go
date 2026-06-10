package slice

// DifferenceInts возвращает элементы из a, которых нет в b. Порядок и повторы из a сохраняются.
func DifferenceInts(a []int, b []int) []int {
	// TODO: реализовать функцию.
	result := make([]int, 0, len(a))
	for _, value := range a {
		found := false
		for _, item := range b {
			if value == item {
				found = true
			}
		}
		if found == false {
			result = append(result, value)
		}
	}
	return result
}
