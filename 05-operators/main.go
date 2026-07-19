package main

import "fmt"

func main() {
	// Arithmetic operators
	A, B, C, D, E, F, G := 10, 20, 11, 3, true, false, 30

	fmt.Println("A: 10 | B: 20 | C: 11 | D: 3 | E: true | F: false | G: 30")

	fmt.Println("A + B =", A+B)
	fmt.Println("A - B =", A-B)
	fmt.Println("A * B =", A*B)
	fmt.Println("A / B =", A/B)
	fmt.Println("C % D =", C%D)

	A++
	fmt.Println("A++ =", A)

	A--
	fmt.Println("A-- =", A)

	// Relational operators
	fmt.Println("A == B?:", A == B)
	fmt.Println("A != B?:", A != B)
	fmt.Println("A >= B?:", A >= B)
	fmt.Println("A <= B?:", A <= B)

	// Logical operators
	fmt.Println("E && F:", E && F)
	fmt.Println("E || F:", E || F)
	fmt.Println("!E:", !E)
	fmt.Println("!F:", !F)

	// Assigment operators
	A += B
	fmt.Println("A += B:", A)

	A -= B
	fmt.Println("A -= B:", A)

	// Bitwise operators (G = 30 = 11110 in binary)
	fmt.Println("G & 2:", G&2)   // AND
	fmt.Println("G ^ 2:", G^2)   // XOR
	fmt.Println("G | 2:", G|2)   // OR
	fmt.Println("~G:", ^G)       // NOT
	fmt.Println("G << 1:", G<<2) // Left shift
	fmt.Println("G >> 1:", G>>2) // Right shift

	// Miscelaneous operators
	ptr := &A
	fmt.Println("Address of A:", ptr)   // Memory address location of A
	fmt.Println("Value of *ptr:", *ptr) // Pointer dereference operator
}
