# go-os-signals [![Go Report Card](https://goreportcard.com/badge/github.com/jamillosantos/go-os-signals)](https://goreportcard.com/report/github.com/jamillosantos/go-os-signals)

## Motivation

Handling signals with go is simple and awesome. But, it has limitations when you
want to implement tests for what should happen when an specific signal arrives
at determined time.

This library tries not to change the Go logic for signals. Neither tries to be
complex. It just wraps the signal behavior into a `interface` that can be mocked.

## Usage

Below a pretty straightforward usage of the library.

```go
package main

import (
	"fmt"
	"syscall"

	signals "github.com/jamillosantos/go-os-signals"
)

func main() {
	l := signals.NewListener(syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Running ... [hit ctrl+c to finish]")
	sig := <-l.Receive()
	fmt.Println("Signal:", sig)
}
```

If you need multiple goroutines to listen to the same signal you can either use multiple `Listener`s or a `Waiter`:

```go
package main

import (
	"fmt"
	"sync"
	"syscall"

	signals "github.com/jamillosantos/go-os-signals"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	l := signals.NewWaiter(syscall.SIGINT, syscall.SIGTERM)
	go func() {
		defer wg.Done()
		l.Wait()
		fmt.Println("From GOROUTINE 1")

	}()
	go func() {
		defer wg.Done()
		l.Wait()
		fmt.Println("From GOROUTINE 2")
	}()

	fmt.Println("Running ... [hit ctrl+c to finish]")
	l.Wait()
	fmt.Println("Signal received")
	wg.Wait() // Wait goroutines to finish
}

```

You can find more examples at the `examples` folder.

## Testing

For testing, the "receiving" part of the code is EXACTLY the same. But, the `MockListener`
has one more capability: `Send`.

```go
package main

import (
	"fmt"
	"syscall"
	"time"

	"github.com/jamillosantos/go-os-signals/signaltest"
)

func main() {
	l := signaltest.NewMockListener(syscall.SIGINT, syscall.SIGTERM)

	go func() {
		time.Sleep(time.Second * 1)
		l.Send(syscall.SIGINT)
	}()

	fmt.Println("Running ... [wait 1 second]")
	sig := <-l.Receive()
	fmt.Println("Signal:", sig)
}
```

You can find more examples at the `examples` folder.
