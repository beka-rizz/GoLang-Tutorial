package basics

import "fmt"

// recover -> built-in function that stops the propagation of a panic and returns the value passed to the panic call

// recover only returns the value, passed to the panic function.


// Use cases: graceful recovery, cleanup, logging and reporting
// Best practices:
// 1) Always use with defer
// 2) Avoid silent recovery
// 3) Avoid overuse

func run16() {
	process()	
	fmt.Println("Returned from process")
}

func process() {
	defer func() {
		// if r := recover(); r != nil {
		r := recover()
		if r != nil {
			fmt.Println("Recovered", r)
		}
	}()

	fmt.Println("Start Process")
	panic("Something went wrong!")
	// fmt.Println("End Process") // will not be executed
}