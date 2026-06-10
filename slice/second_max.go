package slice

import "sort"

// SecondMax возвращает второй по величине уникальный элемент. Если его нет, возвращает 0.
func SecondMax(src []int) int {
	// TODO: реализовать функцию.
	if len(src) == 0 {
		return 0
	}
	copySrc := make([]int, len(src))
	copy(copySrc, src)
	sort.Slice(copySrc, func(i, j int) bool {
		return copySrc[i] > copySrc[j]
	})
	maxValue := copySrc[0]
	for i := 1; i < len(copySrc); i++ {
		if copySrc[i] < maxValue {
			return copySrc[i]
		}
	}
	return 0
}
