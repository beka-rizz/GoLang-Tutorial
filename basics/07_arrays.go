package basics

import "fmt"

func run7() {
	// Array is value type in Go

	// Array declaration
	// var arrayName [size]elementType

	var numbers [5]int // generates array with default values
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println("Fruits array:", fruits)

	// originalArray := [3]int{1, 2, 3}
	// copiedArray := originalArray

	// copiedArray[0] = 100

	// fmt.Println("Original array:", originalArray)
	// fmt.Println("Copied array:", copiedArray)

	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int

	copiedArray = &originalArray
	copiedArray[0] = 100

	fmt.Println("Original array:", originalArray)
	fmt.Println("Copied array:", *copiedArray)
}