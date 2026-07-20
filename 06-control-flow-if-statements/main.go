package main

import "fmt"

func main() {
	fmt.Println("----------------------------------------------------")
	fmt.Println("--- IF STATAMENTS ----------------------------------")
	fmt.Println("----------------------------------------------------")

	age := 17

	if age <= 18 {
		fmt.Println("You cannot vote")
	} else if age >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("Is your age a secret?")
	}
}
