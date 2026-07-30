package channels

func CountChannel(ch <-chan int) int {
	// TODO: посчитать количество значений в канале до закрытия.
	counter := 0
	for range ch {
		counter++
	}
	return counter
}
