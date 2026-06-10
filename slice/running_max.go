package slice

// RunningMax возвращает слайс текущих максимумов при проходе слева направо.
func RunningMax(src []int) []int {
	// TODO: реализовать функцию.
	if len(src) == 0 {
		return src
	}
	result := make([]int, len(src))
	result[0] = src[0]
	localMax := src[0]
	for i := 0; i < len(src); i++ {
		if src[i] > localMax {
			localMax = src[i]
		}
		result[i] = localMax
	}
	return result
}
