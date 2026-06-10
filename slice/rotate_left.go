package slice

// RotateLeft возвращает новый слайс, циклически сдвинутый влево на k позиций. Для k больше длины используется остаток от деления.
func RotateLeft(src []int, k int) []int {
	// TODO: реализовать функцию.
	if len(src) == 0 {
		return []int{}
	}

	k = k % len(src)

	result := make([]int, 0, len(src))
	result = append(result, src[k:]...)
	result = append(result, src[:k]...)
	return result
}
