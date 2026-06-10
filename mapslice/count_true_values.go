package mapslice

// CountTrueValues возвращает количество ключей со значением true.
func CountTrueValues(m map[string]bool) int {
	// TODO: реализовать функцию.
	result := 0
	for _, value := range m {
		if value == true {
			result++
		}
	}
	return result
}
