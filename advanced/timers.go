package main

import (
	"fmt"
	"time"
)

// === MULTIPLE TIMERS
func TimeTest4() {
	timer1 := time.NewTimer(time.Second)
	timer2 := time.NewTimer(2 * time.Second)

	select {
	case <- timer1.C:
		fmt.Println("Timer1 expired")
	case <-timer2.C:
		fmt.Println("Timer2 expired")
	}
}

// === SCHEDULING DELAYED OPERATIONS
func TimeTest3() {
	timer := time.NewTimer(2 * time.Second) // non-blocking timer starts

	go func () {
		<- timer.C
		fmt.Println("Delayed operation executed.")
	}()

	fmt.Println("Waiting...")
	time.Sleep(3 * time.Second) // blocking timer starts
	fmt.Println("End of program")
}

func longRunningOperation() {
	for i := range 20 {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}


// === TIMEOUT
func TimeTest2() {
	timeout := time.After(3 * time.Second)
	done := make(chan bool)

	go func() {
		longRunningOperation()
		done <- true
	}()

	select {
	case <-timeout:
		fmt.Println("Operation timed out.")
	case <- done:
		fmt.Println("Operation completed.")
	}
}

// === BASIC TIMER USE
func TimeTest1() {
	fmt.Println("Starting app.")
	timer := time.NewTimer(2 * time.Second) // non-blocking
	fmt.Println("Waiting for timer.C")
	stopped := timer.Stop()
	if stopped {
		fmt.Println("Timer stopped.")
	}
	timer.Reset(time.Second)
	fmt.Println("Timer reset.")
	<-timer.C // blocking in nature
	fmt.Println("Timer expired.")
}
