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
