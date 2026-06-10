package slice

// PadRight возвращает новый слайс длиной не меньше size, дополняя справа value. Если len(src) >= size, возвращает копию src.
func PadRight(src []int, size int, value int) []int {
	// TODO: реализовать функцию.
	result := make([]int, size)
	copy(result, src)

	if len(src) >= size {
		srcCopy := make([]int, len(src))
		copy(srcCopy, src)
		return srcCopy
	}
	for i := len(src); i < size; i++ {
		result[i] = value
	}
	return result
}
