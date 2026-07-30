package channels

func FanOut(in <-chan int, n int) []<-chan int {
	// TODO: раздать значения из in по n выходным каналам. Каждое значение попадает только в один выход.
	sliceOfCh := make([]<-chan int, n)
	outs := make([]chan int, n)
	for i := 0; i < n; i++ {
		outs[i] = make(chan int, 100)
		sliceOfCh[i] = outs[i]
	}
	go func() {
		index := 0

		for value := range in {
			outs[index] <- value

			index++
			if index == n {
				index = 0
			}
		}

		for _, ch := range outs {
			close(ch)
		}
	}()
	return sliceOfCh
}
