package signals

import (
	"os"
	"sync"
)

// Waiter implements a locking mechanism that can be called multiples, locking multiple goroutines, and unblocks
// all at once when a signal arrives.
//
// Internally it uses the Listener.Receives alongside a multiple
type Waiter interface {
	Wait()
}

// NewWaiter creates a new Waiter initializing the listener with the given signals.
func NewWaiter(signals ...os.Signal) Waiter {
	return &waiter{
		listener: NewListener(signals...),
	}
}

// NewWaiterWithListener creates a new Waiter with the given Listener.
func NewWaiterWithListener(listener Listener) Waiter {
	return &waiter{
		listener: listener,
	}
}

type waiter struct {
	condMutex   sync.Mutex
	cond        *sync.Cond
	startLocker sync.Mutex
	start       sync.Once
	listener    Listener
	wg          sync.WaitGroup
}

func (w *waiter) Wait() {
	w.startLocker.Lock()
	w.start.Do(func() {
		w.cond = sync.NewCond(&w.condMutex)
		go func() {
			<-w.listener.Receive()
			w.cond.Broadcast()

			w.startLocker.Lock()
			w.start = sync.Once{} // Restart the start
			w.startLocker.Unlock()
		}()
	})
	w.startLocker.Unlock()

	w.cond.L.Lock()
	w.cond.Wait()
	w.cond.L.Unlock()
}
