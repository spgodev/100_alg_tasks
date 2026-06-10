package slice

// RemoveAt возвращает новый слайс без элемента по index. Если index вне диапазона, возвращает копию src.
func RemoveAt(src []int, index int) []int {
	// TODO: реализовать функцию.
	if index < 0 || index >= len(src) {
		return append([]int{}, src...)
	}

	result := make([]int, 0, len(src)-1)

	result = append(result, src[:index]...)
	result = append(result, src[index+1:]...)

	return result
}
