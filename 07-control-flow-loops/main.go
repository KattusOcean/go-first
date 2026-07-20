package main

import "fmt"

func main() {
	fmt.Println("----------------------------------------------------")
	fmt.Println("--- FOR LOOP ---------------------------------------")
	fmt.Println("----------------------------------------------------")

	// Basic loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("----------------------------------------------------")
	fmt.Println("--- CONITNUE OR BREAK ------------------------------")
	fmt.Println("----------------------------------------------------")

	// Coninue or break
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("----------------------------------------------------")
	fmt.Println("--- GOTO -------------------------------------------")
	fmt.Println("----------------------------------------------------")

	// GoTo
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			goto end
		}
	}

end:
	fmt.Println("Loop ended")

	// GoTo 2
	i := 0
loop:
	if i < 10 {
		fmt.Println(i)
		i++
		goto loop
	}
}
