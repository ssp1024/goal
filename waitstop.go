package goal

//NewWaitStop create new instance.
func NewWaitStop() *WaitStop {
	return &WaitStop{make(chan struct{})}
}

// WaitStop is lightweight synchronize usage, has two state `not stopped` and `have stoped`.
//
// The initial state is `not stopped`.
type WaitStop struct {
	ch chan struct{}
}

// Stop change state to 'have stopped'.
func (w *WaitStop) Stop() {
	close(w.ch)
}

// Stopped return the state
func (w *WaitStop) Stopped() bool {
	select {
	case <-w.ch:
		return true
	default:
		return false
	}
}

// Wait will block execution until state became to 'have stopped`
func (w *WaitStop) Wait() {
	<-w.ch
}

// C return readonly channel to use in select...case...
//	select {
//		case <-waitStop.C():
//		case <-time.After(time.Second):
//	}
func (w *WaitStop) C() <-chan struct{} {
	return w.ch
}
