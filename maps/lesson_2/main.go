package main

import "fmt"

func main() {
	fruits := make([]string, 3, 6) // can predefine a array size and cap ( to which it could extended) if we know the amount.

	fruits = append(fruits, "mango")
	fruits[0] = "apple"

	fmt.Println(fruits)

	fruits = append(fruits, "orange")
	fruits = append(fruits, "jack fruit")
	fmt.Println(fruits)

	tvs := make(map[string]float64, 3) // can use make with map, but it only takes one parameter unlike cap in array.

	tvs["sony"] = 4.5
	tvs["LG"] = 4.6
	tvs["samsung"] = 5

	fmt.Println(tvs)
}
