package slice

// ContainsValue сообщает, встречается ли target в src.
func ContainsValue(src []int, target int) bool {
	// TODO: реализовать функцию.
	for i := range src {
		if src[i] == target {
			return true
		}
	}
	return false
}
