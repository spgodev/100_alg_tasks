package slice

// InsertAt возвращает новый слайс, где value вставлен по index. Если index вне диапазона [0,len(src)], возвращает копию src.
func InsertAt(src []int, index int, value int) []int {
	// TODO: реализовать функцию.
	result := make([]int, len(src)+1)
	copy(result, src)
	copySrc := make([]int, len(src))
	copy(copySrc, src)
	if index > len(src) || index < 0 {
		return copySrc
	}
	copy(result[:index], src[:index])
	result[index] = value
	copy(result[index+1:], src[index:])
	return result
}
