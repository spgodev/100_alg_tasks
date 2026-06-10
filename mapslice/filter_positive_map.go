package mapslice

// FilterPositiveMap возвращает новую map только с парами, где значение строго положительное.
func FilterPositiveMap(m map[string]int) map[string]int {
	// TODO: реализовать функцию.
	result := make(map[string]int)
	for key, value := range m {
		if value > 0 {
			result[key] = value
		}
	}
	return result
}
