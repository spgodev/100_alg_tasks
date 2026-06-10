package slice

// MatrixRowSums возвращает суммы строк матрицы. Для пустой строки сумма равна 0.
func MatrixRowSums(matrix [][]int) []int {
	// TODO: реализовать функцию.
	result := make([]int, 0, len(matrix))
	for i := range matrix {
		sum := 0
		for j := range matrix[i] {
			sum += matrix[i][j]
		}
		result = append(result, sum)
	}
	return result
}
