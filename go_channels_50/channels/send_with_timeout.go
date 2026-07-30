package channels

import "time"

func SendWithTimeout(ch chan<- int, value int, timeout time.Duration) bool {
	// TODO: попробовать отправить значение в канал с timeout.
	select {
	case ch <- value:
		return true
	case <-time.After(timeout):
		return false
	}
}
