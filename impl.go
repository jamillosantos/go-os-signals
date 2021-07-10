package signals

import (
	"os"
	"os/signal"
)

type listener struct {
	signals []os.Signal
	sigCh   chan os.Signal
}

// NewListener returns the default implementation of a `Listener`.
func NewListener(signals ...os.Signal) Listener {
	ch := make(chan os.Signal)
	signal.Notify(ch, signals...)
	return &listener{
		signals: signals,
		sigCh:   ch,
	}
}

// Receive returns the receive only channel from where the signals will be written to.
func (l *listener) Receive() <-chan os.Signal {
	return l.sigCh
}

// Stop stops listening the signal.
func (l *listener) Stop() {
	close(l.sigCh)
	signal.Stop(l.sigCh)
}
