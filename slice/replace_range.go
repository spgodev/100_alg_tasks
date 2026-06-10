package slice

// ReplaceRange возвращает копию src, где элементы с индексами [from, to) заменены на value. Некорректный диапазон не меняет слайс.
func ReplaceRange(src []int, from int, to int, value int) []int {
	// TODO: реализовать функцию.
	copySrc := make([]int, len(src))
	copy(copySrc, src)
	if from < 0 || to > len(src) || from >= to {
		return copySrc
	}
	for i := from; i < to; i++ {
		copySrc[i] = value
	}
	return copySrc
}
