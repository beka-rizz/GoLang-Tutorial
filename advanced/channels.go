package main

import (
	"fmt"
	"time"
)

// var := make(chan Type)
// Channels are blocking

func Channels() {
	greeting := make(chan string)
	greetString := "Hello"

	// greeting <- greetString // blocking because it continuously trying to reveive values, it is ready to receive continuous flow of data.

	go func() {
		greeting <- greetString
		greeting <- "World"

		for _, e := range "abcde" {
			greeting <- "Alphabet: " + string(e)
		}
	}()
	// go func() {
	// 	receiver := <-greeting
	// 	fmt.Println(receiver)
	// 	receiver = <- greeting
	// 	fmt.Println(receiver)
	// }()

	receiver := <-greeting
	fmt.Println(receiver)
	receiver = <-greeting
	fmt.Println(receiver)

	for range 5 {
		receiver = <- greeting
		fmt.Println(receiver)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("End of program")
}
