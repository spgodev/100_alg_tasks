package mapslice

// RemoveKeys возвращает копию m без ключей из keys.
func RemoveKeys(m map[string]int, keys []string) map[string]int {
	// TODO: реализовать функцию.
	result := make(map[string]int)
	remove := make(map[string]bool)

	for _, key := range keys {
		remove[key] = true
	}
	for key, value := range m {
		if !remove[key] {
			result[key] = value
		}
	}
	return result
}
