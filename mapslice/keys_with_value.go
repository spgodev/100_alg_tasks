package mapslice

import "sort"

// KeysWithValue возвращает ключи, у которых значение равно value, в лексикографическом порядке.
func KeysWithValue(m map[string]int, value int) []string {
	// TODO: реализовать функцию.
	result := make([]string, 0)
	for k, v := range m {
		if v == value {
			result = append(result, k)
		}
	}
	sort.Strings(result)
	return result
}
