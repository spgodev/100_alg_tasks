package slice

// RotateRight возвращает новый слайс, циклически сдвинутый вправо на k позиций.
func RotateRight(src []int, k int) []int {
	// TODO: реализовать функцию.
	if len(src) == 0 {
		return []int{}
	}

	k = k % len(src)

	result := make([]int, 0, len(src))
	result = append(result, src[len(src)-k:]...)
	result = append(result, src[:len(src)-k]...)
	return result
}
