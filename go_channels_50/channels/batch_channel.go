package channels

func BatchChannel(in <-chan int, size int) <-chan []int {
	// TODO: группировать значения из канала батчами размера size.
	out := make(chan []int)
	go func() {
		defer close(out)
		batch := make([]int, 0, size)
		for value := range in {
			batch = append(batch, value)
			if len(batch) == size {
				out <- batch
				batch = make([]int, 0, size)
			}
		}
		if len(batch) != 0 {
			out <- batch
		}
	}()
	return out
}
