package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

// declaring a variable

func main() {

	// fmt.Print("Investment Amount: ")
	// fmt.Scan(&investmentAmount) // take user input

	investmentAmount := investmentAmt("Investment Amount: ")

	// fmt.Print("Expected Return Rate: ")
	// fmt.Scan(&expectedReturnRate)
	expectedReturnRate := expectedReturnRt("Expected Return Rate: ")

	// fmt.Print("Years: ")
	// fmt.Scan(&years)
	years := yr("Years: ")

	// var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// var futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)
	futureValue, futureRealValue := valueCalc(investmentAmount, expectedReturnRate, years)

	// fmt.Println("futureValue: ", futureValue)
	// fmt.Println("futureRealValue: ", futureRealValue)

	prnt(futureValue, futureRealValue)
}

func investmentAmt(msg string) float64 {
	var investmentAmount float64
	fmt.Print(msg)
	fmt.Scan(&investmentAmount)

	return investmentAmount
}

func expectedReturnRt(msg string) float64 {
	var expectedReturnRate float64
	fmt.Print(msg)
	fmt.Scan(&expectedReturnRate)

	return expectedReturnRate
}

func yr(msg string) float64 {
	var years float64
	fmt.Print(msg)
	fmt.Scan(&years)

	return years
}

func valueCalc(investmentAmount, expectedReturnRate, years float64) (float64, float64) {

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)

	return futureValue, futureRealValue

}

func prnt(futureValue, futureRealValue float64) {

	fmt.Printf("futureValue: %0.2f\n", futureValue)
	fmt.Printf("futureRealValue: %0.2f", futureRealValue)

}

// basic structure

// func value(variable1 type, variable2 type .....) (type, type ....) {
// 	...
// 	logic
// 	...
// 	return variable1, variable2

// }
