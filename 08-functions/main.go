package main

import "fmt"

func main() {
	greetings()

	fmt.Println("---------------------------------------------------")

	fmt.Printf("Max value is: %d\n", max(9, 38))

	firstName, lastName := swap("Jhon", "Doe")
	fmt.Printf("Swapped names: %s %s\n", firstName, lastName)
}

func greetings() {
	fmt.Println("Hello, partner!")
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func swap(x, y string) (string, string) {
	return y, x
}
