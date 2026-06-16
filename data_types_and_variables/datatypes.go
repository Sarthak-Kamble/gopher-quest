package main

import "fmt"

func main() {

	// * Data Types in Golang

	// * Integer Types

	// ? uint
	// Platform-dependent unsigned integer type.
	// Size depends on the system architecture:
	// 32-bit system → 32 bits
	// 64-bit system → 64 bits
	// Stores values from 0 upwards.

	// ? uint8
	// Unsigned 8-bit integer.
	// Range: 0 to 255
	// Commonly used for bytes, RGB colors, raw binary data.

	// ? uint16
	// Unsigned 16-bit integer.
	// Range: 0 to 65,535
	// Useful when you need more range than uint8 but less memory than uint32.

	// ? uint32
	// Unsigned 32-bit integer.
	// Range: 0 to 4,294,967,295
	// Commonly used for file sizes, hashes, network protocols.

	// ? uint64
	// Unsigned 64-bit integer.
	// Range: 0 to 18,446,744,073,709,551,615
	// Used for very large numbers like timestamps, IDs, counters.

	// | Type     | Size       | Range               |
	// | -------- | ---------- | ------------------- |
	// | `uint8`  | 8 bits     | 0 → 255             |
	// | `uint16` | 16 bits    | 0 → 65,535          |
	// | `uint32` | 32 bits    | 0 → ~4.29 billion   |
	// | `uint64` | 64 bits    | 0 → ~18 quintillion |
	// | `uint`   | 32/64 bits | Depends on system   |

	var age uint8 = 25

	fmt.Println(age)

	// * Float Types

	// float32
	// 32-bit floating-point number.
	// Uses single precision.
	// Range: approximately ±1.18 × 10⁻³⁸ to ±3.4 × 10³⁸
	// Precision: about 7 decimal digits.
	// Uses less memory but is less accurate.
	// Commonly used when memory matters (graphics, large datasets).

	// float64
	// 64-bit floating-point number.
	// Uses double precision.
	// Range: approximately ±2.23 × 10⁻³⁰⁸ to ±1.8 × 10³⁰⁸
	// Precision: about 15 decimal digits.
	// More accurate and the default choice in Go.
	// Commonly used for calculations, scientific work, finance, etc.

	var price float32 = 99.99
	var distance float64 = 24352.234444

	fmt.Println(price)
	fmt.Println(distance)

	// * byte and rune type in Golang

	// byte
	// Alias for uint8.
	// Represents an 8-bit unsigned integer.
	// Range: 0 to 255.
	// Commonly used for raw bytes, binary data, and ASCII characters.

	// rune
	// Alias for int32.
	// Represents a Unicode code point.
	// Range: -2,147,483,648 to 2,147,483,647.
	// Commonly used for characters because it can store Unicode characters.

	var b byte = 65

	fmt.Println(b)
	fmt.Printf("%c\n", b)

	var r rune = '₹'

	fmt.Println(r)

	
	// string
	// Represents a sequence of characters.
	// Internally stored as a sequence of bytes (UTF-8).
	// Strings are immutable (cannot be changed after creation).
	// Commonly used for text, messages, filenames, JSON, etc.

	text := "Hello Gophers"

	fmt.Println(text)

	// bool
	// Represents a boolean value.
	// Can only store two values:
	// true  → yes/on/condition satisfied
	// false → no/off/condition not satisfied
	// Commonly used in conditions and logic.

	var isLoggedIn bool = true

	fmt.Println("Is Logged In:", isLoggedIn)

	// nil type which is basically null in Golang

	// ! var num uint8 = 288  overflow error

	// fmt.Println(num)


	// * Implicit Assignment

	implicit := "A"   // Majorly used when we initialized the variable

	fmt.Printf("%T", implicit)

	// var number int32

	// number = 10

	// fmt.Println("\n", number)

	fmt.Println("")
	// * Type Casting
	number := uint(10) 
	
	// ? "%T" will return the type
	fmt.Printf("%T", number)
	
	fname := "Sarthak"
	lname := "Kamble"
	designation := "Software Engineer"

	// ? "%v" is basically accessing the value 
	fmt.Printf("\nHello I am %v %v I am a %v\n", fname, lname, designation)

	// ? "%b" is used to display binary representation 
	num2 := 14

	fmt.Printf("%b", num2)

	fmt.Println("")
	// ? "%e" is used to display the scientific notation
	num3 := 14524.24434

	fmt.Printf("%e", num3)


}