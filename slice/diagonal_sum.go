package slice

// DiagonalSum возвращает сумму главной диагонали квадратной или прямоугольной матрицы, пока существуют и строка, и столбец.
func DiagonalSum(matrix [][]int) int {
	// TODO: реализовать функцию.
	i, sum := 0, 0
	for i < len(matrix) && i < len(matrix[i]) {
		sum += matrix[i][i]
		i++
	}
	return sum
}
