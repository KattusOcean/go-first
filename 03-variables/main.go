package main

import "fmt"

func main() {
	//  Static type declaration
	// 1st way: Declare and assign on 2 different lines
	var mango string
	mango = "This is a mango"

	var weight int
	weight = 45

	// 2nd way: Declare and assign on the same line
	var height int = 20

	// Prints
	fmt.Println("Mango:", mango)
	fmt.Println("Weight:", weight)
	fmt.Println("Height:", height)

	// Multiple one-liner declarations
	var apples, trees int = 18, 7

	fmt.Println("Apples:", apples, "| Trees", trees)

	// Dynamic type declaration
	// 3rd way (shorhand -> ":="):
	// GO has "type inference" so there it will
	// determine the data type automatically
	age := 54

	fmt.Println("Age:", age)
	fmt.Printf("Age is of type: %T\n", age)

	// Mixed variable declaration
	a, b, c := 750.52, 8, "Madrid"
	fmt.Printf("A is of type: %T\n", a)
	fmt.Printf("B is of type: %T\n", b)
	fmt.Printf("C is of type: %T\n", c)
}
