package eight

func Channels() {
	ch := make(chan bool, 1)

	ch <- true
	go func() {
		<-ch
	}()
	ch <- true
}
