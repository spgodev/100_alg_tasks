package channels

import "sync"

func MergeTwoChannels(a <-chan int, b <-chan int) <-chan int {
	// TODO: объединить два входных канала в один выходной.
	out := make(chan int, 2)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for v := range a {
			out <- v
		}
	}()
	go func() {
		defer wg.Done()
		for v := range b {
			out <- v
		}
	}()
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
