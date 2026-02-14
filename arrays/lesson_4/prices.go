package main

import "fmt"

func main() {

	prices := []float64{23, 43.0, 11.00}
	fmt.Println(prices)

	// we can append as many as values we want
	prices = append(prices, 45.00, 100) // like this
	fmt.Println(prices)

	// appending values from another list .. required unpacking the list
	morePrices := []float64{45.99, 99, 100}
	// prices = append(prices, morePrices) // this will throw an error, since this is not list of list, it appends on float values.
	prices = append(prices, morePrices...) // ** these three dots tells list to unpack and then append
	fmt.Println(prices)

}
