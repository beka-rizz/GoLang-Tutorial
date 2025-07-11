package main

import (
	"fmt"
	"time"
)

// SELECT -> it allows to wait on multiple channel operations simultaneously and perform corresponding action when one of them becomes ready.
// - Multiplexing channel op-s
// - Non-blocking mechanism
// - Preventing potential deadlocks

// When the channel is closed, select will receive zero value of a channel's type
// We must handle this scenario via boolean value that comes with received value

func main() {
	ChannelCheck()
}

func ChannelCheck() {
	ch := make(chan int)

	go func() {
		ch <- 1
		close(ch) // signal that no more values will be sent on that channel
	}()

	for {
		select {
		case msg, isClosed := <-ch:
			if !isClosed {
				fmt.Println("Channel closed")
				// clean up activities
				return
			}
			fmt.Println("Received:", msg)
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout.")
		}
	}
}

func SelectWithTimeout() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		close(ch)
	}()

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout.")
		// case timee := <-time.After(time.Second):
		// 	fmt.Println("Timeout.")
		// 	fmt.Println("Current time:", timee)
	}
}

func Select2() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
	}()

	go func() {
		time.Sleep(time.Second)
		ch2 <- 2
	}()

	time.Sleep(2 * time.Second) // to be sure that all channels are received values

	for range 2 { // we loop N times per N channel (we want all received values from channels)
		select {
		case msg := <-ch1:
			fmt.Println("Received from ch1:", msg)
		case msg := <-ch2:
			fmt.Println("Received from ch2:", msg)
		default:
			fmt.Println("No channels ready...")
		}
	}
	fmt.Println("End of program")
}

func Select1() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	for i := range 10 {
		go func() {
			fmt.Println("no timer goroutine in loop", i+1)
			ch1 <- 1
		}()

		go func() {
			time.Sleep(3 * time.Second)
			fmt.Println("3 second goroutine in loop", i+1)
			ch2 <- 2
		}()

		time.Sleep(500 * time.Millisecond)
		select {
		case msg := <-ch1:
			fmt.Println("Received from ch1:", msg)
		case msg := <-ch2:
			fmt.Println("Received from ch2:", msg)
		default:
			fmt.Println("No channels ready...")
		}
		fmt.Println("End of loop:", i+1, "-----------")
	}
	fmt.Println("End of program")
}
