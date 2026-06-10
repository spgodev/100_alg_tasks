package slice

// RemoveDuplicatesSorted для отсортированного src удаляет соседние дубликаты и возвращает уникальные элементы.
func RemoveDuplicatesSorted(src []int) []int {
	// TODO: реализовать функцию.
	if len(src) == 0 {
		return src
	}
	remover := 1
	for i := 1; i < len(src); i++ {
		if src[i] != src[remover-1] {
			src[remover] = src[i]
			remover++
		}
	}
	src = src[:remover]
	return src
}
