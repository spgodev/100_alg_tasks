package channels

func SumChannel(ch <-chan int) int {
	// TODO: посчитать сумму всех чисел из канала до его закрытия.
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}
