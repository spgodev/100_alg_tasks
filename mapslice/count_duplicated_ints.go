package mapslice

// CountDuplicatedInts возвращает количество различных значений, которые встречаются больше одного раза.
func CountDuplicatedInts(src []int) int {
	// TODO: реализовать функцию.
	counts := make(map[int]int)
	result := 0
	for _, value := range src {
		counts[value]++
	}
	for _, v := range counts {
		if v != 1 {
			result++
		}
	}
	return result
}
