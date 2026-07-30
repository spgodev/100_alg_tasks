package channels

import "sync"

func LimitConcurrency(tasks []func(), limit int) {
	// TODO: выполнить список задач, но одновременно не больше limit.
	wg := sync.WaitGroup{}
	sem := make(chan struct{}, limit)
	for _, task := range tasks {
		sem <- struct{}{}
		wg.Add(1)
		go func(task func()) {
			defer wg.Done()
			defer func() {
				<-sem
			}()

			task()
		}(task)
	}
	wg.Wait()
}
