package basics

import "fmt"

// Defer => defers the execution of a function until the surrounding function returns

// Defer function is always enclosed with another function

// The deferred call's arguments are evaluated immediately

// Note: defer function executes regardless of whether surrounding function returns normally or due to an error


// Use cases: resource cleanup, unlocking mutexes, logging and tracing
// Best practices:
// 1) Keep deferred actions short
// 2) Understand evaluation timing
// 3) Avoid complex control flow

func run14() {
	process()
	process2(10)
}

func process() {
	defer fmt.Println("First deferred statement executed.")
	defer fmt.Println("Second deferred statement executed.")
	defer fmt.Println("Third deferred statement executed.")
	fmt.Println("This is normal execution statement.")
	fmt.Println("End of process function.")
}

func process2(i int) {
	defer fmt.Println("Deferred i value:", i)
	i++
	fmt.Println("Value of i:", i)
}