package basics

import "fmt"

// Function in go -> First Class Object (First Class Citizen)

func run11() {
	// func <name>(parameter list) returnType {
	// code block
	// return value
	// }

	// fmt.Println(add(3, 4))

	// func(){
	// 	fmt.Println("This is Anonymous Function")
	// }() // anonymous function


	// greet := func(){
	// 	fmt.Println("Hello from Anonymous Function")
	// }
	
	// greet()

	// operation := add
	// result := operation(5, 22)
	// fmt.Println(result)
	
	// Passing a func as an argument
	result := applyOperation(5, 3, add)
	fmt.Println(result)

	// Returning and using a function
	muptiplyBy2 := createMultiplier(2)
	fmt.Println("6 * 2 =", muptiplyBy2(6))

}

func add(a, b int) int {
	return a + b
}

// Function that takes a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// Function that returns a function
func createMultiplier(factor int) func(int) int {
	return func (x int) int {
		return x * factor
	}
}