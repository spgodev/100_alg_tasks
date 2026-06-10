package slice

// TakeEverySecond возвращает элементы с индексами 0, 2, 4 и так далее.
func TakeEverySecond(src []int) []int {
	// TODO: реализовать функцию.
	result := make([]int, 0, len(src)/2)
	for i, value := range src {
		if i%2 == 0 {
			result = append(result, value)
		}

	}
	return result
}
