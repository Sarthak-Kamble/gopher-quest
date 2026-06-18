package main

import "fmt"

func main() {

	// * For Loop in Golang

	fmt.Println("------------------ For Loop in Golang ------------------")

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Number: ", i)
	// }

	str := "hello Gophers✨"

	fmt.Println(str[0]) // This output Integer representation of h
	fmt.Println(string(str[0]))

	// ? ASCII -> 1 Byte 256
	// ? UTF-8 -> 4 Bytes

	for idx := 0; idx < len(str); idx++ {
		fmt.Printf("%c", str[idx])
	}

	// Adding _ instead of i as I don't require index variable so you can make use of underscore
	for _, char := range str {
		fmt.Printf("%c", char)
	}


	// * Exercise

	// number := 14

	// Write a program which will print a table of the number from 1 to 12

	// for idx := 1; idx <= 12; idx++ {
	// 	fmt.Printf("%v x %v = %v\n", number, idx, number * idx)
	// }



}
