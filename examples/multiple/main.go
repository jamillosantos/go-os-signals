package main

import (
	"fmt"
	"syscall"

	signals "github.com/setare/go-os-signals"
)

func main() {
	l2 := signals.NewListener(syscall.SIGINT, syscall.SIGTERM)
	l := signals.NewListener(syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-l2.Receive()
		fmt.Println("From GOROUTINE: Signal:", sig)
	}()

	fmt.Println("Running ... [hit ctrl+c to finish]")
	sig := <-l.Receive()
	fmt.Println("Signal:", sig)
}
