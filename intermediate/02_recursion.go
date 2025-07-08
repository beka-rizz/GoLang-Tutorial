package intermediate

import "fmt"

// Use cases: mathematical algos, tree and graph traversal, divide and conquer algos

func run2() {
	fmt.Println(factorial(10))
	fmt.Println(sumOfDigits(10))
	fmt.Println(sumOfDigits(19123))
	fmt.Println(fibonacci(3))
	fmt.Println(fibonacci(4))
	fmt.Println(fibonacci(5))
	fmt.Println(fibonacci(20))
}

func fibonacci(n int) int {
	if n == 1 {
		return 0
	} 
	if n == 2 {
		return 1
	}
	return fibonacci(n - 1) + fibonacci(n - 2)
}

func sumOfDigits(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}

func factorial(n int) int {
	// Base Case => 0! = 1
	if n == 0 {
		return 1
	}
	return n * factorial(n - 1)
}