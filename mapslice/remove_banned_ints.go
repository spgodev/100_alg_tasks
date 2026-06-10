package mapslice

// RemoveBannedInts возвращает элементы src, которых нет в banned или для которых banned[value] == false.
func RemoveBannedInts(src []int, banned map[int]bool) []int {
	// TODO: реализовать функцию.
	result := make([]int, 0)
	for i := range src {
		if !banned[src[i]] {
			result = append(result, src[i])
		}
	}
	return result
}
