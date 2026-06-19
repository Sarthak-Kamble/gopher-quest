package main

import "fmt"

func main() {

	// var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// for idx := 1; idx < len(arr); idx++ {
	// 	if arr[idx]%2 == 0 {
	// 		fmt.Println(arr[idx])
	// 	}
	// }

	// skills := [5]string {"JavaScript", "TypeScript", "Golang", "React", "TailwindCSS"}

	// fmt.Println("-------- Skills --------")
	// for _, skill := range skills {
	// 	fmt.Println(skill)
	// }

	// matrix := [2][2]int {{1, 2}, {3, 4}}

	// fmt.Println(matrix)

	// matrix := [...][2]int {{1, 2}, {3, 4}, {5, 6}}

	// fmt.Println(matrix)
	
	// printMutatedMatrix(matrix)

	// fmt.Println(matrix)

	// * Array Exercises

	programming_lang := [5]string {"C", "C++", "Golang", "Rust", "Zig"}

	// fmt.Println(programming_lang)

	for _, lang := range programming_lang {
		fmt.Println(lang)
	}

	numbers := [5]int{10,20,30,40,50}

	number_set := [6]int{4,9,2,15,7,1}

	nums := [7]int{1,2,3,4,5,6,8}
	
	var sum int = 0

	for _, num := range numbers {
		sum += num
	}

	// fmt.Println(sum)
	// * Let's print this part in a formatted way
	fmt.Printf("Sum of numbers is: %v", sum)

	fmt.Printf("\nAverage of numbers is: %v", sum / len(numbers))

	var max_num int = number_set[0] 

	for _, num := range number_set {
		if num > max_num {
			max_num = num
		}
	}

	fmt.Printf("\nMaximum number in number_set is: %v", max_num)


	// Count even numbers in the "nums" array

	// Initializing a variable to count even numbers
	var even_count int = 0

	for _, num := range nums {
		if num % 2 == 0 {
			even_count++
		}
	}

	fmt.Printf("\nCount of even numbers in 'nums' %v", even_count)

}

// func printMutatedMatrix (arr [3][2]int) {

// 	arr[0] = [2]int {10, 10}

// 	fmt.Println(arr)

// }