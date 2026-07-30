package channels

import "time"

func TimeoutResult(fn func() int, timeout time.Duration) (int, bool) {
	// TODO: запустить fn, вернуть результат или false при timeout.
	resultCh := make(chan int, 1)
	go func() {
		resultCh <- fn()
	}()
	select {
	case result := <-resultCh:
		return result, true
	case <-time.After(timeout):
		return 0, false

	}
}
