package slice

// UniqueInts возвращает уникальные значения в порядке первого появления.
func UniqueInts(src []int) []int {
	// TODO: реализовать функцию.
	result := make([]int, 0, len(src))
	for _, value := range src {
		isSeen := false
		for i := range result {
			if result[i] == value {
				isSeen = true
			}
		}
		if !isSeen {
			result = append(result, value)
		}
	}
	return result
}
