package nine

import (
	"context"
	"sync"
)

func merge(channels ...chan int) chan int {
	var wg sync.WaitGroup
	mergedChannel := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())

	out := func(ch chan int) {
		defer wg.Done()

		for {
			select {
			case num, ok := <-ch:
				if !ok {
					cancel()
					return
				}

				select {
				case mergedChannel <- num:
				case <-ctx.Done():
					return
				}

			case <-ctx.Done():
				return
			}
		}
	}

	for _, ch := range channels {
		wg.Add(1)
		go out(ch)
	}

	go func() {
		wg.Wait()
		close(mergedChannel)
	}()

	return mergedChannel
}
