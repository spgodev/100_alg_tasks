package slice

// MoveZerosEnd возвращает новый слайс, где все нули перенесены в конец, а порядок ненулевых элементов сохранён.
func MoveZerosEnd(src []int) []int {
	// TODO: реализовать функцию.
	result := make([]int, 0, len(src))
	zeroCounter := 0
	for i := range src {
		if src[i] != 0 {
			result = append(result, src[i])
		}
		if src[i] == 0 {
			zeroCounter++
		}
	}
	for i := 0; i < zeroCounter; i++ {
		result = append(result, 0)
	}
	return result
}
