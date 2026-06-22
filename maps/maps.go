package main

import "fmt"

func main() {

	// * Maps in Golang

	// Declaring a Map
	// mp := map[string]int{"a": 1, "b": 2}
	mp := map[string][]int{"a": {1, 2, 3}}

	fmt.Println(mp)
	
	mp["b"] = []int {4, 5, 6}
	
	// fmt.Println(mp)
	
	fmt.Println(mp["a"])
	fmt.Println(mp["b"])
	
	// Delete a key in the map using the delete func
	// delete(mp, "b")
	// fmt.Println(mp)
	
	// Checking if a key exists in the map
	
	// value, ok := mp["b"] // Returns [4, 5, 6] true as this exist in our map
	value, ok := mp["c"]  // Returns [] false as this key doesn't exist in our map

	fmt.Println(value, ok)
	
	// Exercise

	skills := map[string][]string{"Frontend": {"JavaScript", "React", "TailwindCSS", "TypeScript"}}

	fmt.Println(skills)
	
	skills["Backend"] = []string {"ExpressJS", "Fastify", "NodeJS", "PostgreSQL"}
	fmt.Println(skills)

}