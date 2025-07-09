package main

import (
	"fmt"
	"time"
)

// =========== SYNCHRONIZING DATA EXCHANGE
func Sync4() {
	data := make(chan string)

	go func() {
		for i := range 5 {
			data <- fmt.Sprintf("hello %d", i)
			time.Sleep(100 * time.Millisecond)
		}
		close(data)
	}()
	// close(data) // Channel is closed before Goroutine could send a value to the channel

	for value := range data {
		fmt.Println("Received value:", value, ":", time.Now())
	} // loops over only on active channel, creates receiver each time and stops creating receiver (looping) once the channel is closed
}

// =========== SYNCHRONIZING MULTIPLE GOROUTINES AND ENSURING THAT ALL GOROUTINES ARE COMPLETE
func Sync3() {
	numGoroutines := 3
	done := make(chan int, 3)

	for i := range numGoroutines {
		go func(id int) {
			fmt.Printf("Goroutine %d working...\n", id)
			time.Sleep(time.Second)
			done <- id // SENDING SIGNAL OF COMPLETION
		}(i)
	}

	for range numGoroutines {
		<-done // Wait for each goroutine to finish, WAIT FOR ALL GOROUTINES TO SIGNAL COMPLETION
	}

	fmt.Println("All goroutines are finished.")
}

func Sync2() {
	ch := make(chan int)

	go func() {
		ch <- 9 // blocking until the value is received
		time.Sleep(1 * time.Second)
		fmt.Println("Sent value")
	}()

	value := <-ch // blocking until the value is sent
	fmt.Println("Value:", value)
}

func Sync1() {
	done := make(chan struct{})

	go func() {
		fmt.Println("Working...")
		time.Sleep(3 * time.Second)
		done <- struct{}{}
	}()
	<-done
	fmt.Println("Finished.")
}
