package mapslice

// BuildIntSet строит set из src в виде map[int]bool, где каждое встреченное значение имеет true.
func BuildIntSet(src []int) map[int]bool {
	// TODO: реализовать функцию.
	result := make(map[int]bool)
	for _, value := range src {
		result[value] = true
	}
	return result
}
