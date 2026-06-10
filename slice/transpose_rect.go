package slice

// TransposeRect транспонирует прямоугольную матрицу. Для nil или пустой матрицы возвращает пустой слайс.
func TransposeRect(matrix [][]int) [][]int {
	// TODO: реализовать функцию.
	if len(matrix) == 0 {
		return [][]int{}
	}
	rows := len(matrix)
	cols := len(matrix[0])

	transposed := make([][]int, cols)
	for i := range transposed {
		transposed[i] = make([]int, rows)
	}
	for i := range matrix {
		for j := range matrix[i] {
			transposed[j][i] = matrix[i][j]
		}
	}
	return transposed
}
