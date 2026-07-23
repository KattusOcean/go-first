package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a String
	greeting := "Hello, World!"
	fmt.Println("Normal string: " + greeting)

	// Print the Hexadecimal byte values of a String
	fmt.Print("Hexadecimal bytes: ")
	for i := 0; i < len(greeting); i++ {
		fmt.Printf("%x", greeting[i])
	}
	fmt.Println("")

	// Measuring String lenght
	fmt.Printf("String lenght is: %d", len(greeting))
	fmt.Println("")

	// String concatenation (slice of a string -> concatenation with Join)
	fruits := []string{"apples", "oranges", "and bananas"}
	fmt.Println(strings.Join(fruits, " "))
}
