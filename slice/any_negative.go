package slice

// AnyNegative возвращает true, если в src есть хотя бы одно отрицательное число.
func AnyNegative(src []int) bool {
	// TODO: реализовать функцию.
	for i := range src {
		if src[i] < 0 {
			return true
		}
	}
	return false
}
