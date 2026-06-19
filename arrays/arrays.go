package main

import "fmt"

func main() {

	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for idx := 1; idx < len(arr); idx++ {
		if arr[idx]%2 == 0 {
			fmt.Println(arr[idx])
		}
	}

	skills := [5]string {"JavaScript", "TypeScript", "Golang", "React", "TailwindCSS"}

	fmt.Println("-------- Skills --------")
	for _, skill := range skills {
		fmt.Println(skill)
	}

	// matrix := [2][2]int {{1, 2}, {3, 4}}

	// fmt.Println(matrix)

	matrix := [...][2]int {{1, 2}, {3, 4}, {5, 6}}

	fmt.Println(matrix)
	
	printMutatedMatrix(matrix)

	fmt.Println(matrix)
}

func printMutatedMatrix (arr [3][2]int) {

	arr[0] = [2]int {10, 10}

	fmt.Println(arr)

}