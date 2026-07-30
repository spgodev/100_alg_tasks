package channels

func CloseChannel(ch chan int) {
	// TODO: закрыть переданный канал.
	close(ch)
}
