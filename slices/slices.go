package main

import "fmt"

func main() {

	// * Slices in Golang
	arr := [5]int{1, 2, 3, 4, 5}

	sl := arr[:3]

	fmt.Println(sl)

}