package main

import (
	"fmt"
	"time"
)

// Goroutines are just functions that leave the main thread and run in the background and come back to join the main thread once the functions are finished/to return any value.
// Goroutines do not stop the program flow and are non blocking
// Goroutines run independently and concurrently

func Rungo() {
	var err error
	fmt.Println("Beginning program.")
	go sayHello()
	fmt.Println("After sayHello function.")

	go func() {
		err = doWork()
	}()
	// err = go doWork() // This is not accepted
	go printNumbers()
	go printLetters()

	time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("doWork completed successfully!")
	}
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine")
}

func printNumbers() {
	for i := range 5 {
		fmt.Println("Number:", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println(string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func Test() {
	name := "Beka"
	fmt.Println("Name before goroutine:", name)
	// go func(str *string) {
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println(*str)
	// }(&name)
	go func(str string) {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}(name)
	name = "Ais"
	fmt.Println("Name is:", name)
	time.Sleep(2 * time.Second)
}

func doWork() error {
	// Simulate work
	time.Sleep(1 * time.Second)
	return fmt.Errorf("an error occured in doWork")
}
