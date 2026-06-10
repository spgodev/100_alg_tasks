package slice

// ReverseCopy возвращает новый слайс с элементами src в обратном порядке.
func ReverseCopy(src []int) []int {
	// TODO: реализовать функцию.
	result := make([]int, len(src))
	copy(result, src)
	n := len(result)
	for i := 0; i < n/2; i++ {
		result[i], result[n-1-i] = result[n-1-i], result[i]
	}
	return result
}
