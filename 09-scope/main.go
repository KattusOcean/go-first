package main

import "fmt"

// 2. Global variables
var g int

// 3. Shadowed variables
var x int = 25

func main() {
	// 1. Local variables
	var a, b int = 10, 20
	fmt.Println("A:", a)

	// 2. Global variables
	g = a + b
	fmt.Println("G:", g)

	// 3. Shadowed variables (local variable shadows global var)
	var x int = 100
	fmt.Println("X:", x)
}
