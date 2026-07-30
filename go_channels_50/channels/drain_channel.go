package channels

func DrainChannel(ch <-chan int) {
	// TODO: прочитать все значения из канала и ничего не возвращать.

	for range ch {
	}
}
