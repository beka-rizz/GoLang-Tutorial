package basics

import (
	"fmt"
	"os"
)

// os.exit -> will not invoke deferred functions, including those registered using defer

func run17() {
	defer fmt.Println("Deferred statement")
	fmt.Println("Starting the main function")
	
	// Exit with status code of 1
	os.Exit(1)
	
	// This will never be executed
	fmt.Println("End of main function")
}