package signals

import "os"

// Listener abstracts the behavior of a `signal.Notify` and `signal.Reset`.
//
// The idea is to `signal.Notify` will be called at the initialization at the
// implementation.
type Listener interface {
	// Receive returns the receive only channel from where the signals will be
	// written to.
	Receive() <-chan os.Signal

	// Stop stops listening the signal.
	Stop()
}
