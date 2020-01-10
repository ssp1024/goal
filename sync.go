package goal

import (
	"sync"
)

//SelectWaitGroup replace sync.WaitGroup with channel which can used in select...case...
//	wg := otherFn()
//	select {
//		case <-SelectWaitGroup(wg):
//		case <-time.After(time.Second):
//	}
func SelectWaitGroup(wg *sync.WaitGroup) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}
