package slice

// FlattenMatrix разворачивает матрицу в один слайс построчно.
func FlattenMatrix(matrix [][]int) []int {
	// TODO: реализовать функцию.
	result := make([]int, 0)
	for i := range matrix {
		for j := range matrix[i] {
			result = append(result, matrix[i][j])
		}
	}
	return result
}
