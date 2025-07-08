package basics

import "fmt"

// Variadic functions -> defined by prefixing the type of the last param with an ellipsis.

// ... <=> Ellipsis

// Syntax:
// func functionName(param1 type1, param2 type2, param3 ...type3) {}

// variadic params must be at the END of the param list when we have regular params too.

// Destructuring slice: => sliceName...

func run13() {
	statement, total := sum("The sum of 1, 2, 3 is", 1, 2, 3)
	fmt.Println(statement, total)
	
	sequence, total := sum2(1, 10, 20, 30, 40, 50)
	fmt.Printf("Sequence: %d, total: %d\n", sequence, total)
	
	numbers := []int{1, 2, 3, 4, 5, 9}
	
	sequence2, total2 := sum2(2, numbers...)
	fmt.Printf("Sequence2: %d, total2: %d\n", sequence2, total2)

}

func sum(returnString string, nums ...int) (string, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return returnString, total
}

func sum2(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}