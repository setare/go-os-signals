package signals

import (
	"os"
)

// MockListener is a `Listener` that allows the developer to fake signals for
// testing purposes.
type MockListener interface {
	Listener
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

// Send channels the signal to the listener.
func (l *mockListener) Send(s os.Signal) {
	l.sigCh <- s
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
