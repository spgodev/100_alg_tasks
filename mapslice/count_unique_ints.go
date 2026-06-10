package mapslice

// CountUniqueInts возвращает количество уникальных значений в src.
func CountUniqueInts(src []int) int {
	// TODO: реализовать функцию.
	seen := make(map[int]int)
	result := 0
	for _, value := range src {
		seen[value]++
	}
	result = len(seen)

	return result
}
