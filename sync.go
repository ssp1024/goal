package goal

import (
	"sync"
)

//SelectWaitGroup replace sync.WaitGroup with readonly channel, compatible with select...case...
func SelectWaitGroup(wg *sync.WaitGroup) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}
