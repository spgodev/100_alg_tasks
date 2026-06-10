package mapslice

// IntFrequency возвращает map значение -> количество вхождений.
func IntFrequency(src []int) map[int]int {
	// TODO: реализовать функцию.
	result := make(map[int]int)
	for _, value := range src {
		result[value]++
	}
	return result
}
