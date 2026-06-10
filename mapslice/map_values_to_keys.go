package mapslice

import "sort"

// MapValuesToKeys возвращает map значение -> слайс ключей с этим значением, ключи внутри каждого слайса должны быть отсортированы лексикографически.
func MapValuesToKeys(m map[string]int) map[int][]string {
	// TODO: реализовать функцию.
	result := make(map[int][]string)

	for k, v := range m {
		result[v] = append(result[v], k)
	}
	for v := range result {
		sort.Strings(result[v])
	}
	return result
}
