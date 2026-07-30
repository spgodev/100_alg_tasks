package channels

import "time"

func ReceiveWithTimeout(ch <-chan int, timeout time.Duration) (int, bool) {
	// TODO: ждать значение из канала не дольше timeout.
	select {
	case result := <-ch:
		return result, true
	case <-time.After(timeout):
		return 0, false
	}
}
