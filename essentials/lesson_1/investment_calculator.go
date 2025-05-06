package main

import (
	"fmt"
	"math"
)

// passive approach

// func main() {

// 	var investmentAmount = 1000
// 	var expectedReturnRate = 5.6
// 	var years = 20

// 	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))

// 	fmt.Println(futureValue)
// }

// active approach

func main() {
	const inflationRate = 2.5
	var investmentAmount float64 // declaring a variable
	var expectedReturnRate float64
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount) // take user input

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

// short code

// func main() {

// 	investmentAmount, expectedReturnRate, years := 1000.0, 5.6, 20.0

// 	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

// 	fmt.Println(futureValue)
// }
