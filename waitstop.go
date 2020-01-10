package goal

func NewWaitStop() *WaitStop {
	return &WaitStop{make(chan struct{})}
}

type WaitStop struct {
	ch chan struct{}
}

func (w *WaitStop) Stop() {
	close(w.ch)
}

func (w *WaitStop) Stopped() bool {
	select {
	case <-w.ch:
		return true
	default:
		return false
	}
}

func (w *WaitStop) Wait() {
	<-w.ch
}

func (w *WaitStop) C() <-chan struct{} {
	return w.ch
}
