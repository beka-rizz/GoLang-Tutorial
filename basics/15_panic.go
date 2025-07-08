package basics

import "fmt"

// panic -> built-in function that stops the ordinary flow of control and begins panicking

// when panic is called, the current function stops execution, and any deferred functions in that function are executed, and then control returns to the calling function

func run15() {
	process(-1)
}

func process(input int) {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0 {
		fmt.Println("Before Panic")
		panic("input must be a non-negative number")
	}
	fmt.Println("Processing input:", input)
}