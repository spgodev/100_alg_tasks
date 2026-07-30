package channels

import "time"

func RunWithTimeout(fn func(), timeout time.Duration) bool {
	// TODO: запустить fn и вернуть true, если она завершилась до timeout.
	done := make(chan struct{}, 1)
	go func() {
		fn()
		done <- struct{}{}
	}()
	select {
	case <-done:
		return true
	case <-time.After(timeout):
		return false
	}
}
