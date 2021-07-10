package main

import (
	"fmt"
	"syscall"

	signals "github.com/jamillosantos/go-os-signals"
)

func main() {
	c := make(chan bool, 1)

	l2 := signals.NewListener(syscall.SIGINT, syscall.SIGTERM)
	l := signals.NewListener(syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-l2.Receive()
		fmt.Println("From GOROUTINE: Signal from l2:", sig)
		close(c)
	}()

	fmt.Println("Running ... [hit ctrl+c to finish]")
	sig := <-l.Receive()
	fmt.Println("Signal from l:", sig)

	<-c // Wait the goroutine to close
}
