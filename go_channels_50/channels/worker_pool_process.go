package channels

import "sync"

func WorkerPoolProcess(in <-chan int, workers int, fn func(int) int) <-chan int {
	// TODO: обработать все числа функцией fn параллельно через workers.
	outs := make(chan int)
	wg := sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for value := range in {
				outs <- fn(value)
			}
		}()

	}
	go func() {
		wg.Wait()
		close(outs)
	}()
	return outs
}
