package basics

import (
	"fmt"
	"slices"
)

func run9() {
	// var sliceName[]ElementType

	// var numbers []int
	// var numbers1 = []int{1, 2, 3}

	// numbers2 := []int{1, 2, 3}

	// slice1 := make([]int, 5)

	a := [5]int{1, 2, 3, 4, 5}

	slice1 := a[:2]
	slice1 = a[0:]
	slice1 = a[3:4]
	fmt.Println(slice1)
	
	slice1 = append(slice1, 6)
	fmt.Println("Slice1:", slice1)
	
	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)

	fmt.Println("Slice1 copy:", sliceCopy)

	// var nilSlice []int

	for i, v := range slice1 {
		fmt.Printf("Index: %d, value: %d\n", i, v)
	}

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("Two slices are equal!")
	}
}