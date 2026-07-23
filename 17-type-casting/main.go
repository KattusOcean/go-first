package main

import (
	"fmt"
	"strconv"
)

func main() {
	var total int = 20
	var items int = 7
	var average float64

	// This is the `int result = (int) Math.round(...)`
	average = float64(total) / float64(items)
	fmt.Printf("Average = %.2f\n", average)

	// String to integer
	str := "123"
	num, err := strconv.Atoi(str) // Atoi = Ascii to integer
	if err != nil {
		fmt.Println("Error convertirng:", err)
		return
	}

	fmt.Printf("Conerted number: %v\n", num)
}
