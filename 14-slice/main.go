package main

import "fmt"

func main() {
	// Slices are like ArrayLists in Java
	// Their syntax and functionality is basically an array
	// but they offer more flexibility with its size

	// Note: the length is the current number of elements
	// that are stored in the slice, the capacity is the
	// total amount of elements that the slice can hold

	var numbers []int        // Declaration of a slice of integers
	numbers = make([]int, 5) // Inizialization of the slice with a size and length of 5 ([0, 0, 0, 0, 0])
	fmt.Println("Slice is:", numbers)

	// Subslicing
	numbers2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8} // Slice declaration (shorthand variation)
	subSlice1 := numbers2[1:4]                   // Will have numbers from index 1 to 4, last one not included ([1,2,3])
	subSlice2 := numbers2[:3]                    // Will have numbers from the starting index to the third ([0,1,2])
	subSlice3 := numbers2[4:]                    // Will have numbers from the fourth index to the last one ([3,4,5,6,7,8])

	fmt.Println("Subslice1:", subSlice1)
	fmt.Println("Subslice2:", subSlice2)
	fmt.Println("Subslice3:", subSlice3)

	// Append() and Copy() methods
	numbers3 := []int{}
	numbers3 = append(numbers3, 0, 1, 2, 3, 4)

	numbers4 := make([]int, len(numbers3), cap(numbers3)*2) // len() == length ; cap() == capacity
	copy(numbers4, numbers3)

	fmt.Println("Numbers3 slice:", numbers3)
	fmt.Println("Numbers4 slice after copy:", numbers4)

	// Nil slices
	var numbers5 []int
	fmt.Println("Nil slice:", numbers5)
}
