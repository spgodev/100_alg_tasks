package slice

// IntersectInts возвращает уникальные элементы, которые есть и в a, и в b, в порядке первого появления в a.
func IntersectInts(a []int, b []int) []int {
	// TODO: реализовать функцию.
	result := make([]int, 0, len(a))
	for i := range a {
		value := a[i]
		existInB := false
		alreadyAdded := false
		for j := range b {
			if value == b[j] {
				existInB = true
			}
			for k := range result {
				if value == result[k] {
					alreadyAdded = true
				}
			}
		}
		if existInB && !alreadyAdded {
			result = append(result, value)
		}
	}
	return result
}
