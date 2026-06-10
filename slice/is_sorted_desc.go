package slice

// IsSortedDesc сообщает, отсортирован ли src по невозрастанию.
func IsSortedDesc(src []int) bool {
	// TODO: реализовать функцию.
	if len(src) == 0 {
		return true
	}

	for i := 0; i < len(src)-1; i++ {
		if src[i+1] > src[i] {
			return false
		}
	}
	return true
}
