package slice

// DropEveryNth возвращает слайс без каждого n-го элемента при счёте с 1. Если n <= 0, возвращает копию src.
func DropEveryNth(src []int, n int) []int {
	// TODO: реализовать функцию.
	result := make([]int, 0, len(src))
	if n <= 0 {
		copyResult := make([]int, len(src))
		copy(copyResult, src)
		return copyResult
	}
	for i := range src {
		if (i+1)%n == 0 {
			continue
		}
		result = append(result, src[i])
	}
	return result
}
