package slice

// CountEven возвращает количество чётных чисел в src. Для nil или пустого слайса возвращает 0.
func CountEven(src []int) int {
	// TODO: реализовать функцию.
	if len(src) == 0 {
		return 0
	}
	counter := 0
	for i := range src {
		if src[i]%2 == 0 {
			counter++
		}
	}
	return counter
}
