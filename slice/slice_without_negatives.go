package slice

// SliceWithoutNegatives возвращает новый слайс без отрицательных чисел.
func SliceWithoutNegatives(src []int) []int {
	// TODO: реализовать функцию.
	result := make([]int, 0, len(src))
	for _, value := range src {
		if value >= 0 {
			result = append(result, value)
		}
	}
	return result
}
