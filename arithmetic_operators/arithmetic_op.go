package main

import "fmt"

func main() {

	// Arithmetic Operators in Golang
	x := uint8(10)
	y := 290

	// z := x + uint8(y)  // ! Overflow when we convert var y to uint8

	z := int(x) + y

	a := 25
	b := 5

	fmt.Println(z)
	
	div := a / b
	fmt.Println(div)

	d := 20
	e := 2

	product := d * e

	fmt.Println(product)

	// * Exercises

	// ? Calculate Area of Rectangle

	length := 20
	breadth := 5

	area_of_rect := length * breadth

	fmt.Printf("Area of Rectange: %v\n", area_of_rect)

	// ? Convert Seconds

	seconds := 3672

	hours := seconds / 3600
	minutes := seconds / 60
	remaining_seconds := seconds % 60

	fmt.Println("\n----------- Convert Seconds -----------")
	fmt.Println("Hours: ", hours)
	fmt.Println("Minutes: ", minutes)
	fmt.Println("Remaining Seconds: ", remaining_seconds)
	fmt.Println("-----------------------------------------")
	
	// Swap two numbers without a third variable
	
	num1 := 25
	num2 := 15
	
	fmt.Println("\n----------- Swap Two Numbers -----------")
	
	fmt.Printf("Before Swap: num1 = %v and num2 = %v", num1, num2)
	
	// Swap without a third variable	
	num1 = num1 + num2    // num1 = 25 + 15 -> 40
	num2 = num1 - num2    // num2 = 40 - 15 -> 25
	num1 = num1 - num2    // num1 = 40 - 25 -> 15
	
	fmt.Printf("\nBefore Swap: num1 = %v and num2 = %v\n", num1, num2)

	fmt.Println("------------------------------------------")

}