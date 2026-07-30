package channels

func WaitDone(done <-chan struct{}) {
	// TODO: ждать закрытия канала done.
	<-done
}
