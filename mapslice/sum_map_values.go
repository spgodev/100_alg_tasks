package mapslice

// SumMapValues возвращает сумму всех значений map. Для nil или пустой map возвращает 0.
func SumMapValues(m map[string]int) int {
	// TODO: реализовать функцию.
	result := 0
	for _, v := range m {
		result += v
	}
	return result
}
