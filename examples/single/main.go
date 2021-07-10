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
