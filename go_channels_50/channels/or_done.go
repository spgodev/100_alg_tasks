package channels

func OrDone(done <-chan struct{}, in <-chan int) <-chan int {
	// TODO: читать значения из in, пока не закрыт done.
	return nil
}
