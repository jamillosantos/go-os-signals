package signaltest

import (
	"os"

	"github.com/jamillosantos/go-os-signals"
)

// MockListener is a `Listener` that allows the developer to fake signals for
// testing purposes.
type MockListener interface {
	signals.Listener
	Send(os.Signal)
}

type mockListener struct {
	signals []os.Signal
	sigCh   chan os.Signal
}

// NewMockListener returns the an implementation of a `MockListener`. This instance
// allows the programmer to send fake signals for testing.
func NewMockListener(signals ...os.Signal) MockListener {
	ch := make(chan os.Signal)
	return &mockListener{
		signals: signals,
		sigCh:   ch,
	}
}

// Send channels the signal to the listener. If `s` is not included at the `signals`
// passed when the instance was created (`NewMockListener`), the signal will not
// be passed to the `Receive` method.
//
// Note that no actual signal is sent.
func (l *mockListener) Send(s os.Signal) {
	for _, ss := range l.signals {
		if ss == s {
			l.sigCh <- s
		}
	}
}

// Receive returns the receive only channel from where the signals will be written
// to.
func (l *mockListener) Receive() <-chan os.Signal {
	return l.sigCh
}

// Stop stops listening the signal.
func (l *mockListener) Stop() {
	close(l.sigCh)
}
