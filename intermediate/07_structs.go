package intermediate

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Address
	PhoneHomeCell // anonymous field
}

type Address struct {
	city 	string
	country string
}

type PhoneHomeCell struct {
	home string
	cell string
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}

func run7() {


	p1 := Person{
		firstName: "Bekarys",
		lastName:  "Nurgasym",
		age:       21,
		address: Address{
			city: "Almaty",
			country: "Kazakhstan",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "1242342",
			cell: "534563212",
		},
		
	}

	p2 := Person{
		firstName: "Aisaule",
		age:       26,
	}

	p2.address.city = "New York"
	p2.address.country = "USA"

	fmt.Println(p1.firstName)
	fmt.Println(p2.firstName)
	fmt.Println(p1.fullName())
	fmt.Println(p1.address)
	fmt.Println(p1.address.country)
	fmt.Println(p2.address.country)
	fmt.Println(p1.cell) // accessing embedded anonymous fields directly
	fmt.Println(p1.home)

	// Anonymous Structs
	// user := struct {
	// 	username string
	// 	email string
	// } {
	// 	username: "user123",
	// 	email: "user123@mail.ru",
	// }
	// fmt.Println(user.username)

	// fmt.Println("Before increment:", p1.age)
	// p1.incrementAgeByOne()
	// fmt.Println("After increment:", p1.age)

}



