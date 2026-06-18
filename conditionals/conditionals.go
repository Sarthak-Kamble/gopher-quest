package main

import "fmt"

func main() {

	fmt.Println("------ Comparison Operators ------")

	a := 10
	b := 20

	// Equal to ==
	fmt.Println("a == b:", a == b)

	// Not equal !=
	fmt.Println("a != b:", a != b)

	// Greater than >
	fmt.Println("a > b:", a > b)

	// Less than <
	fmt.Println("a < b:", a < b)

	// Greater than or equal >=
	fmt.Println("a >= 10:", a >= 10)

	// Less than or equal <=
	fmt.Println("b <= 20:", b <= 20)


	fmt.Println("\n------ Type Conversion Example ------")

	x := 8
	y := uint8(10)

	// x < y gives error because:
	// x is int
	// y is uint8

	result := uint8(x) < y

	fmt.Println("uint8(x) < y:", result)


	fmt.Println("\n------ Using Comparison with if-else ------")

	age := 20

	if age >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You cannot vote")
	}


	fmt.Println("\n------ Even Odd Example ------")

	number := 7

	if number%2 == 0 {
		fmt.Println(number, "is Even")
	} else {
		fmt.Println(number, "is Odd")
	}


	fmt.Println("\n------ Multiple Conditions ------")

	marks := 75

	if marks >= 40 && marks <= 100 {
		fmt.Println("Student Passed")
	} else {
		fmt.Println("Student Failed")
	}


	fmt.Println("\n------ Login Example ------")

	username := "admin"
	password := "12345"

	if username == "admin" && password == "12345" {
		fmt.Println("Login Successful")
	} else {
		fmt.Println("Invalid Credentials")
	}


	fmt.Println("\n------ NOT Operator ------")

	isLoggedIn := false

	if !isLoggedIn {
		fmt.Println("Please login first")
	}


	// * Switch Case in Golang

	fmt.Println("----------------- Switch Case in Golang -----------------")
	
	day := "Monday"
	
	switch day {
	case "Sunday":
		fmt.Println("Weekend")
	case "Saturday":
		fmt.Println("Weekend")
	case "Monday":
		fmt.Println("Weekday")
	case "Tuesday":
		fmt.Println("Weekday")
	default:
		fmt.Println("Please tell me its weekend")
	}
	
	// * Naked Switch statement in Golang
	fmt.Println("\n----------------- Naked Switch statement in Golang -----------------")

	my_age := 25

	switch {
	case my_age < 18:
		fmt.Println("You are a Teenager")
		fallthrough
	case my_age >= 18:
		fmt.Println("You are an Adult")
		fallthrough
	default:
		fmt.Println("You are old")
	}
	
	

}