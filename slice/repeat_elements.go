package slice

// RepeatElements повторяет каждый элемент src times раз подряд. Если times <= 0, возвращает пустой слайс.
func RepeatElements(src []int, times int) []int {
	// TODO: реализовать функцию.
	if times <= 0 {
		return []int{}
	}
	result := make([]int, 0, len(src)*times)
	for i := range src {
		for j := 0; j < times; j++ {
			result = append(result, src[i])
		}
	}
	return result
}
