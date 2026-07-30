package channels

func Broadcast(in <-chan int, n int) []<-chan int {
	// TODO: каждое значение из in отправить во все n выходных каналов.
	sliceOfCh := make([]<-chan int, n)
	outs := make([]chan int, n)
	for i := 0; i < n; i++ {
		outs[i] = make(chan int, 100)
		sliceOfCh[i] = outs[i]
	}
	go func() {

		for value := range in {
			for i := range outs {
				outs[i] <- value
			}

		}

		for _, ch := range outs {
			close(ch)
		}
	}()
	return sliceOfCh
}
