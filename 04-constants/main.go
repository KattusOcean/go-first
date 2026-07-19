package main

import "fmt"

func main() {
	// Declare general constants
	// A good practice is to uppercase constant variable names
	const AGE int = 19
	const CITY string = "Valladolid"
	const PI float64 = 3.14
	const ISACTIVE bool = true

	fmt.Println("Age:", AGE)
	fmt.Println("City", CITY)
	fmt.Println("Pi:", PI)
	fmt.Println("Is active?:", ISACTIVE)

	// Declare integer literals
	const (
		DECIMAL = 255  // decimals with no prefix
		OCTAL   = 0377 // octal with leading 0
		HEX     = 0xff // hexadecimal with leading 0x
	)

	fmt.Println("Decimal:", DECIMAL, "| Octal:", OCTAL, "| Hexadecimal:", HEX)

	// Declare floating-point literals
	// A floating-point literal can have an integer part,
	// a decimal point, a fraactional part, and a exponent
	// part
	const AVOGADRO = 6.022e23
	fmt.Println("Avogadro:", AVOGADRO)

	// Escape sequences in string literals
	// - New line
	const GREETING string = "Hello, Moon!\n"

	// - Quote
	const QUOTE string = "/GO is fun!/ - Someone"

	// - Alert or Bell
	const BELL string = "Bell is \a"

	// - Line breaks
	const LB = "I\nAM\nBATMAN\n!"

	fmt.Println(GREETING)
	fmt.Println(QUOTE)
	fmt.Println(BELL)
	fmt.Println(LB)
}
