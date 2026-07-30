package channels

func TrySend(ch chan<- int, value int) bool {
	// TODO: попробовать отправить значение без блокировки.
	select {
	case ch <- value:
		return true
	default:
		return false
	}
}
