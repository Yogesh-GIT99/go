package main

import "fmt"

type device map[string]float64 // type alias: more cleaner code

func (d device) output() {
	fmt.Println(d)
}

func main() {
	fruits := make([]string, 3, 6) // can predefine a array size and cap ( to which it could extended) if we know the amount.

	fruits = append(fruits, "mango")
	fruits[0] = "apple"

	fmt.Println(fruits)

	fruits = append(fruits, "orange")
	fruits = append(fruits, "jack fruit")
	fmt.Println(fruits)

	tvs := make(device, 3) // can use make with map, but it only takes one parameter unlike cap in array.

	tvs["sony"] = 4.5
	tvs["LG"] = 4.6
	tvs["samsung"] = 5

	tvs.output()

	// for loop with array/slices

	// for range fruits  ( can use like this if we donot care about index and values seperately)
	for index, value := range fruits {

		// any other logic ..
		fmt.Println("Index: ", index)
		fmt.Println("Value: ", value)
	}

	// for loop with map

	for key, value := range tvs {

		// any other logic ..
		fmt.Println("Key: ", key)
		fmt.Println("Value: ", value)
	}

}
