package mapslice

// AreDisjoint возвращает true, если у двух слайсов нет общих значений.
func AreDisjoint(a []int, b []int) bool {
	// TODO: реализовать функцию.
	isSeen := make(map[int]bool)
	for _, v := range a {
		isSeen[v] = true
	}
	for _, v := range b {
		if isSeen[v] {
			return false
		}
	}
	return true
}
