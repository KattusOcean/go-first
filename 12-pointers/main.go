package main

import "fmt"

func main() {
	// Storing the address: `&` gets the in memory address of a varibale
	// Dereferencing: `*` access the value at the pointer's address
	var a int = 20
	var ip *int = &a

	fmt.Printf("Address of a: %x\n", &a)
	fmt.Printf("Address stored in ip: %x\n", ip)
	fmt.Printf("Value at *ip: %d\n", *ip)

	// Nil pointers: returns 0 when pointer was not initialized
	var ptr *int
	fmt.Printf("The value of a non initialized pointer is: %x\n", ptr)
	// fmt.Printf("The value of a non initialized pointer is: %x\n", *ptr) - This will throw and error
	fmt.Printf("The value of a non initialized pointer is: %x\n", &ptr)
}
