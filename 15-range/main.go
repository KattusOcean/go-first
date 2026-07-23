package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, number := range numbers {
		fmt.Printf("Slice item %d is %d\n", i, number)
	}

	numbers2 := []int{99, 98, 97}
	for i := range numbers2 {
		fmt.Printf("Array item %d is %d\n", i, numbers2[i])
	}
}
