package basics

import "fmt"

// init -> invoked automatically when the package is initialized

// executes before main

// Use cases: setup tasks, configuration, registering components, db initialization
// Best practices:
// 1) Avoid side effects
// 2) Initilization order
// 3) Documentation

func init() {
	fmt.Println("Initializing package1...")
}

func init() {
	fmt.Println("Initializing package2...")
}

func init() {
	fmt.Println("Initializing package3...")
}

func main() {
	fmt.Println("Insed the main function")

}