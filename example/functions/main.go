package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 2.5
	var investmentAmount float64 // declaring a variable
	var expectedReturnRate float64
	var years float64

	// fmt.Print("Investment Amount: ")
	printmsg("Investment Amount: ")
	fmt.Scan(&investmentAmount) // take user input

	// fmt.Print("Expected Return Rate: ")
	printmsg("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	var f = futureValue(investmentAmount, expectedReturnRate, years)
	var frv = futureRealValue(f, inflationRate, years)

	fmt.Println("futureValue: ", f)
	fmt.Println("futureRealValue: ", frv)
}

func printmsg(x string) {

	fmt.Print(x)
}

func futureValue(investmentAmount float64, expectedReturnRate float64, years float64) float64 {

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	return futureValue
}

func futureRealValue(f float64, inflationRate float64, years float64) float64 {

	var FRV = f / math.Pow(1+inflationRate/100, years)

	return FRV

}

// basic structure

// func value(variable1 type, variable2 type .....) (type, type ....) {
// 	...
// 	logic
// 	...
// 	return variable1, variable2

// }
