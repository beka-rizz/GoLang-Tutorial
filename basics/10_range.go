package basics

import "fmt"

// range -> operates on a COPY of the data structure (DS) it iterates over. 
// therefore, modifying an index or a value inside the loop -> does not affect the original DS.

func run10() {
	message := "Hello World"

	for i, v := range message {
		// fmt.Println(i, v)
		fmt.Printf("Index: %d, Rune: %c\n", i, v)

	}
}