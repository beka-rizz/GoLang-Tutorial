package basics

import (
	"errors"
	"fmt"
)

func run12() {
	
	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d, remainder: %d\n", q, r)

	result, err := compare(3, 2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

}

func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func divide2(a, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which is greater")
	}
}