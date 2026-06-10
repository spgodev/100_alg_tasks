package slice

// CountOdd возвращает количество нечётных чисел в src. Отрицательные нечётные тоже учитываются.
func CountOdd(src []int) int {
	// TODO: реализовать функцию.
	count := 0
	for i := range src {
		if src[i]%2 == 1 || src[i]%2 == -1 {
			count++
		}
	}
	return count
}
