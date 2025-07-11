package main

import (
	"fmt"
	"time"
)

// Unbuffered channels look for an immediate receiver.
// When we use ubf chs, we need sender and receiver ready at the same time.

// We always send values into ubf ch inside a goroutine, and then we receive those values inside the main go routine (main thread), or any other function that we have.

// Receiver channel blocks the code until it receives a value

func UnbufferedRun() {
	// ch := make(chan int)
	// go func() {
	// 	// ch <- 1
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("2 second goroutine finished")
	// }()
	// go func() {
	// 	// ch <- 1
	// 	time.Sleep(3 * time.Second)
	// 	fmt.Println("3 second goroutine finished")
	// }()

	// receiver := <-ch
	// fmt.Println(receiver)
	// fmt.Println("End of program")

	// Test1()
	// Test2()
	// Test3()
	// Test4()
	Test5()
}

func Test1() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	receiver := <-ch
	fmt.Println(receiver)
}

func Test2() { // gives an error, because receiver is not ready
	ch := make(chan int)
	go func() {
		fmt.Println("Goroutine before sending started...")
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine before sending finished.")
	}()
	ch <- 1 // looks for immediate receiver, blocks here (waits for all goroutines before to be finished)
	go func() {
		receiver := <-ch
		fmt.Println(receiver)
	}()
	fmt.Println("End of program")
}

func Test3() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	go func() {
		receiver := <-ch
		fmt.Println(receiver)
	}()
	time.Sleep(300 * time.Millisecond)
	fmt.Println("End of program")
}

func Test4() {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("3 second goroutine finished")
	}()
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println(<-ch)
		fmt.Println("5 second goroutine finished")
	}()
	ch <- 1 // blocks here
	fmt.Println("End of program")
}

func Test5() {
	ch := make(chan int)
	go func() {
		fmt.Println("Sender goroutine started...")
		ch <- 1
		fmt.Println("Sender goroutine finished")
	}()
	go func() {
		fmt.Println("3 second receiver goroutine started...")
		time.Sleep(3 * time.Second)
		fmt.Println(<-ch)
		fmt.Println("3 second receiver goroutine finished")
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("End of program")
}
