package channels

func ReceiveOne(ch <-chan int) int {
	// TODO: прочитать одно число из канала и вернуть его.
	return <-ch
}
