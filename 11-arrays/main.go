package main

import "fmt"

func main() {
	// Declaration of an Array
	var balance1 [10]float32

	// Initialization of an Array at declaration
	balance2 := [5]float32{15.4, 84.20, 2.3, 0.3, 18.94}

	// Initialization leeting GO find the elements
	balance3 := [...]float32{15.4, 84.20, 2.3, 0.3, 18.94}

	for i := 0; i < len(balance3); i++ {
		fmt.Println("Element %d = %f", i, balance3[i])
	}
}
