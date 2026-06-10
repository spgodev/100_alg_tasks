package mapslice

// IncrementMapValues возвращает новую map, где к каждому значению прибавлен delta.
func IncrementMapValues(m map[string]int, delta int) map[string]int {
	// TODO: реализовать функцию.
	result := make(map[string]int)
	for k, v := range m {
		result[k] = v + delta
	}
	return result
}
