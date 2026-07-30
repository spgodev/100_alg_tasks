package channels

func SendOne(ch chan<- int, value int) {
	// TODO: отправить одно число в канал.
	ch <- value
}
