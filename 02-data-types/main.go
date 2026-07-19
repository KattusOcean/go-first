package main

import "fmt"

func main() {
	// 1. Integers
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("--- INTEGER NUMBERS -------------------------------------------")
	fmt.Println("---------------------------------------------------------------")

	// 1.1. Signed integers (can be positive or negative)
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	i = -128
	i8 = 127
	i16 = -32768
	i32 = 2147483647
	i64 = -9223372036854775808

	// 1.2. Unsigned integers (can only be positive)
	var u uint
	var u8 uint8
	var u16 uint16
	var u32 uint32
	var u64 uint64

	u = 255
	u8 = 255
	u16 = 65535
	u32 = 4294967295
	u64 = 18446744073709551615

	// Prints integers
	fmt.Println("Signed integers:", i, i8, i16, i32, i64)
	fmt.Println("Unsigned integers:", u, u8, u16, u32, u64)

	// 2. Float
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("--- FLOAT NUMBERS ---------------------------------------------")
	fmt.Println("---------------------------------------------------------------")

	// 2.1. Float32 (used for less precise calculations)
	var f32 float32 = 10.6

	// 2.2. Float64 (used for more precise calculations)
	var f64 float64 = 10.6

	// Prints floats
	fmt.Println("Float32:", f32)
	fmt.Println("Float64", f64)

	// Demonstation on how precise both of them are
	// HP = High Precision | LP = Low Precision
	var HP float64 = 10123456789012345
	var LP float32 = 10123456789012345

	fmt.Println("High precision float:", HP)
	fmt.Println("Low precision float:", LP)

	// 3. Booleans
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("--- BOOLEANS --------------------------------------------------")
	fmt.Println("---------------------------------------------------------------")

	var isActive bool = true
	var isOn bool = false

	fmt.Println("Is active:", isActive)
	fmt.Println("Is on:", isOn)

	// 4. Complex numbers (real number, imaginary number)
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("--- COMPLEX NUMBERS -------------------------------------------")
	fmt.Println("---------------------------------------------------------------")

	var CN1 complex128 = complex(5, 10)
	var CN2 complex64 = complex(2, 7)

	fmt.Println("CN1:", CN1)
	fmt.Println("CN2:", CN2)

	// 5. Strings
	fmt.Println("---------------------------------------------------------------")
	fmt.Println("--- STRINGS ---------------------------------------------------")
	fmt.Println("---------------------------------------------------------------")

	var name string = "Kattus da Ocean!"
	fmt.Println("Name:", name)
}
