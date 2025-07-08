package intermediate

import "fmt"

func run6() {

	// Printing Functions
	fmt.Print("Hello ")
	fmt.Print("World!")
	fmt.Print(12, 456)
	
	fmt.Println("Hello ")
	fmt.Println("World!")
	fmt.Println(12, 456)
	
	name := "Beka"
	age := 21
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Binary: %b, Hex: %X\n", age, age)
	

	// Formatting Functions
	s := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(s)

	// Scanning Functions
	// var name string
	// var age int

	// fmt.Print("Enter your name and age:")
	fmt.Scan(&name, &age)
	fmt.Scanln(&name, &age)
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	
}