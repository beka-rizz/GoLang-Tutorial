package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func run4() {
	message := "Hello, \nGo!"
	message1 := "Hello, \tGo!"
	message2 := "Hello, \rGo!" // carriage return
	rawMessage := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	// array of unicode chars
	fmt.Println("The length of message var is", len(message))
	fmt.Println("The first character in message var is", message[0])

	greeting := "Hello "
	name := "Alice"
	fmt.Println(greeting + name)

	str1 := "Apple" // a has a ASCII code = 65
	str2 := "Banana" // b has a ASCII code = 98
	str3 := "app" // a has a ASCII code = 97
	fmt.Println(str1 < str2) // lexicographical comparison -> dictionary ordering
	fmt.Println(str3 < str1)

	for _, char := range "Bekarys" {
		// fmt.Printf("Character at index %d is %c\n", i, char)
		fmt.Printf("HEC: %x, ", char)
		fmt.Printf("ASCII: %v => ", char)
		fmt.Printf("%c\n", char)
	}

	fmt.Println("Rune Count:", utf8.RuneCountInString(greeting))

	// strings are immutable

	var ch rune = 'a'
	qch := 'ө'
	fmt.Println(ch)
	fmt.Println(qch)

	// Runes store Unicode values, it occupies 4 byte (int32) of memory	

	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", qch)

	cstr := string(ch)
	fmt.Printf("The type of cstr is %T\n", cstr)

	const BYE = "Қайырлы түн!"
	fmt.Println(BYE)

	jhello := "こんにちは" // Japanese Hello
	for _, runeValue := range jhello {
		fmt.Printf("%c\n", runeValue)
	}

	r := '😊'
	fmt.Printf("%v\n", r)
	fmt.Printf("%c\n", r)
}