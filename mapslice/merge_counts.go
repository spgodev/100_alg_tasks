package mapslice

// MergeCounts объединяет две map со счётчиками, складывая значения одинаковых ключей.
func MergeCounts(a map[string]int, b map[string]int) map[string]int {
	// TODO: реализовать функцию.
	result := make(map[string]int)
	for k, v := range a {
		result[k] = v
	}
	for k, v := range b {
		result[k] += v
	}
	return result
}
