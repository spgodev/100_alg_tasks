package slice

// PrefixSums возвращает слайс префиксных сумм: элемент i равен сумме src[0:i+1].
func PrefixSums(src []int) []int {
	// TODO: реализовать функцию.
	result := make([]int, len(src))
	sum := 0

	for i := range src {
		sum += src[i]
		result[i] = sum
	}

	return result
}
