package mapslice

// ToggleIntSet возвращает копию set, последовательно переключая значения: если value был true — становится false/удаляется, иначе становится true.
func ToggleIntSet(set map[int]bool, values []int) map[int]bool {
	// TODO: реализовать функцию.
	result := make(map[int]bool)
	for k, v := range set {
		result[k] = v
	}
	for _, value := range values {
		if result[value] {
			delete(result, value)
		} else {
			result[value] = true
		}
	}
	return result
}
