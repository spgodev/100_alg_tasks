package mapslice

// IsSubset возвращает true, если каждое уникальное значение subset встречается в set.
func IsSubset(subset []int, set []int) bool {
	// TODO: реализовать функцию.
	if len(subset) == 0 {
		return true
	}
	values := make(map[int]bool)
	for _, value := range set {
		values[value] = true
	}
	for _, value := range subset {
		if !values[value] {
			return false
		}
	}
	return true
}
