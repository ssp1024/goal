package goal

import (
	"sync"
)

func SelectWaitGroup(wg *sync.WaitGroup) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}
