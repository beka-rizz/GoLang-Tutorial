package main

import (
	"fmt"
	"time"
)

func BfTest4() {
	ch := make(chan int, 5)
	f1 := func() {
		for i := 1; i <= 5; i++ {
			go func() {
				time.Sleep(1 * time.Second)
				fmt.Println("Value of i:", i)
				ch <- i
			}()
		}
	}
	f1()
	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)
}

func BfTest3() {

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Receiving from buffer")
	go func() {
		fmt.Println("Goroutine 2 second timer started.")
		time.Sleep(2 * time.Second)
		fmt.Println("Received:", <-ch) // ends <- starts (execution starts from right to left)
	}()
	// fmt.Println("Blocking starts...")
	ch <- 3 // blocks because the buffer is full
	// fmt.Println("Blocking ends.")
	// fmt.Println("Received:", <-ch)
	// fmt.Println("Received:", <-ch)
}

func BfTest2() {
	// ===================== BLOCKING ON RECEIVE ONLY IF THE BUFFER IS EMPTY
	ch := make(chan int, 2)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		time.Sleep(3 * time.Second)
		ch <- 2
	}()
	fmt.Println("First blocking starts...")
	fmt.Println("Value:", <-ch)
	fmt.Println("First blocking ends.")
	fmt.Println("Second blocking starts...")
	fmt.Println("Value:", <-ch)
	fmt.Println("End of program")
}

func BfTest1() {
	// ===================== BLOCKING ON SEND ONLY IF THE BUFFER IS FULL
	// variable := make(chan Type, capacity)

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Received:", <-ch)
	}()
	fmt.Println("Blocking starts...")
	ch <- 3 // blocks because the buffer is full
	fmt.Println("Blocking ends.")
	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)

	fmt.Println("Buffered channels")
}
