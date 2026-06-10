package mapslice

// SetIntersectionStrings возвращает уникальные строки, которые есть в обоих слайсах, в порядке первого появления в a.
func SetIntersectionStrings(a []string, b []string) []string {
	// TODO: реализовать функцию.
	inB := make(map[string]bool)
	added := make(map[string]bool)
	result := make([]string, 0)

	for _, value := range b {
		inB[value] = true
	}

	for _, value := range a {
		if inB[value] && !added[value] {
			result = append(result, value)
			added[value] = true
		}
	}

	return result
}
