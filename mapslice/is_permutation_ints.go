package mapslice

// IsPermutationInts возвращает true, если b является перестановкой a с учётом количества повторов.
func IsPermutationInts(a []int, b []int) bool {
	// TODO: реализовать функцию.
	if len(a) != len(b) {
		return false
	}
	counts := make(map[int]int)
	for _, val := range a {
		counts[val]++
	}
	for _, val := range b {
		if counts[val] == 0 {
			return false
		}
		counts[val]--
	}
	return true
}
