package slice

// IsSortedAsc сообщает, отсортирован ли src по неубыванию. Nil и пустой слайс считаются отсортированными.
func IsSortedAsc(src []int) bool {
	// TODO: реализовать функцию.
	if len(src) == 0 {
		return true
	}

	for i := 1; i < len(src); i++ {
		if src[i-1] > src[i] {
			return false
		}
	}
	return true
}
