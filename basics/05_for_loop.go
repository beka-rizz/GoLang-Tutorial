package basics

import "fmt"

func run5() {
	for i := 0; i < 5; i++ {
		fmt.Println("Line", i)
	}

	numbers := [] int {1,2,3, 4 , 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, value: %d\n", index, value)
	}

	// For as in Python or JS
	for i := range 10 {
		fmt.Println(i)
	}

	// For as while
	num := 10
	for num > 0 {
		fmt.Println("Hello World")
		num--;
	}

	for {
		fmt.Println("Beka is good guy")
		break
	}

}