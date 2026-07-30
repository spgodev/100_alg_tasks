package recursion

// CountDigits рекурсивно считает количество цифр в числе.
// CountDigits(0) = 1
// CountDigits(12345) = 5
// Для отрицательных чисел считать цифры без минуса.
func CountDigits(n int) int {
	// TODO
	if n < 0 {
		n = n * -1
	}
	if n < 10 {
		return 1
	}
	return 1 + CountDigits(n/10)
}
