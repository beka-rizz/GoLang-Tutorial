package main

import (
	"fmt"
	"time"
)

// We use channel declarations only in functions or goroutines

// We must declare bidirectional channel, and pass it as an argument to the func-s or goroutines

func main() {
	ch := make(chan int)

	Producer(ch)
	Consumer(ch)
}

// Send only channel
func Producer(ch chan<- int) {
	go func() {
		for i := range 5 {
			ch <- i
			time.Sleep(400 * time.Millisecond)
		}
		close(ch)
	}()
}

// Receive only channel
func Consumer(ch <-chan int) {
	for val := range ch {
		fmt.Println("Value:", val)
	}
}