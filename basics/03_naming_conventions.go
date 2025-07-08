package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func run3() {
	// PascalCase
	// Eg. CalculateArea, UserInfo, NewHTTPRequest
	// Structs, interface, enums

	// snake_case
	// Eg. user_id, first_name, http_request

	// UPPERCASE
	// Constants

	// misedCase
	// Eg. javaScript, htmlDocument, isValid

	const MAX_RETRIES = 5
	var employee_id = 1001
	fmt.Println("EmployeeID: ", employee_id)
}