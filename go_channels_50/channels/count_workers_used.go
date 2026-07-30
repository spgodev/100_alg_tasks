package channels

import "sync"

func CountWorkersUsed(tasks <-chan int, workers int) int {
	// TODO: запустить worker pool и вернуть количество реально обработанных задач.
	counter := 0
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(workers)
	for w := 0; w < workers; w++ {
		go func() {
			defer wg.Done()
			for range tasks {
				mx.Lock()
				counter++
				mx.Unlock()
			}
		}()
	}
	wg.Wait()
	return counter
}
