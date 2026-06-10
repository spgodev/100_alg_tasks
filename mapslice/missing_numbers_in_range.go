package mapslice

// MissingNumbersInRange возвращает числа из диапазона [left,right], которых нет в src, по возрастанию. Если left > right, возвращает пустой слайс.
func MissingNumbersInRange(src []int, left int, right int) []int {
	// TODO: реализовать функцию.
	if left > right {
		return []int{}
	}
	seen := make(map[int]bool)

	for _, value := range src {
		seen[value] = true
	}

	result := make([]int, 0)

	for num := left; num <= right; num++ {
		if !seen[num] {
			result = append(result, num)
		}
	}
	return result
}
