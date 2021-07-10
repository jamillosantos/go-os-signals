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
