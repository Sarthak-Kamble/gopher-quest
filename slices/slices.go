package main

import "fmt"

func main() {

	// * Slices in Golang
	// arr := []int{1, 2, 3, 4, 5}

	// fmt.Println(len(arr)) // Length -> 5
	
	// Let's find the Data Type of the arr
	// fmt.Printf("%T\n", arr)
	
	// // ? Append Method in Slices
	
	// arr = append(arr, 6, 7, 8)
	// fmt.Println(arr)
	
	// fmt.Println("Length: ", len(arr)) 

	// ? make function

	// skills := []string{}   // ? Normal Declaration

	skills := make([]string, 3, 5) // ? make(type, length, capacity)

	skills = append(skills, "C")
	skills = append(skills, "C++")
	skills = append(skills, "Golang")
	skills = append(skills, "JavaScript")
	skills = append(skills, "TypeScript")


	fmt.Println("Slice: ", skills)
	fmt.Println("Length: ", len(skills))
	fmt.Println("Capacity: ", cap(skills))

}