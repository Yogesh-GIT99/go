package main

import (
	"fmt"
)

func main() {
	hobbies := [3]string{"dancing", "drummer", "musician"}

	fmt.Println(hobbies)      // complete array
	fmt.Println(hobbies[0])   // first element
	fmt.Println(hobbies[1:])  // second and third element in list
	fmt.Println(hobbies[:2])  // first and second element slice
	fmt.Println(hobbies[0:2]) // another way for above solution

	// even after slicing the array remain intact.
	reslice := hobbies[0:2]
	reslice_again := reslice[1:3]
	fmt.Println(reslice_again)

	// dynamic array:

	var goal []string = []string{"learn golang", "learn webdevelopment"}
	fmt.Println(goal)
	goal[1] = "learn web development"
	goal = append(goal, "improve dsa")

	fmt.Println(goal)

	// struct with array

	type product struct {
		title string
		id    int
		price float64
	}

	list := []product{{"soap", 1, 23.0}, {"shampoo", 2, 300.0}}
	item := product{"jug", 3, 40.0}
	list = append(list, item)
	fmt.Println(list)

}
