package slice

// TrimZerosEdges удаляет нули только с начала и конца слайса. Нули внутри сохраняются.
func TrimZerosEdges(src []int) []int {
	// TODO: реализовать функцию.
	left, right := 0, len(src)-1
	for left < len(src) && src[left] == 0 {
		left++
	}
	for right >= 0 && src[right] == 0 {
		right--
	}
	if left > right {
		return []int{}
	}
	return src[left : right+1]
}
