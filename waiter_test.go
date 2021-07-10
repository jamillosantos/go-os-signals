package signals_test

import (
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	signals "github.com/jamillosantos/go-os-signals"
	"github.com/jamillosantos/go-os-signals/signaltest"
)

func Test_waiter_Wait(t *testing.T) {
	var wg sync.WaitGroup

	listener := signaltest.NewMockListener(os.Interrupt)

	goRoutinesFeedback := make([]bool, 10)
	wg.Add(len(goRoutinesFeedback))

	l := signals.NewWaiterWithListener(listener)
	for idx, _ := range goRoutinesFeedback {
		go func(idx int) {
			defer wg.Done()
			l.Wait()
			goRoutinesFeedback[idx] = true
		}(idx)
	}

	go func() {
		time.Sleep(time.Millisecond * 50)
		listener.Send(os.Interrupt)
	}()

	wg.Wait() // Wait goroutines to finish

	for idx, feedback := range goRoutinesFeedback {
		require.True(t, feedback, "failed checking goroutine #%d", idx)
	}
}
