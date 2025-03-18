package main

import "fmt"

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64
	outputText("revenue: ")
	fmt.Scan(&revenue)

	outputText("expenses: ")
	fmt.Scan(&expenses)

	outputText("taxRate: ")
	fmt.Scan(&taxRate)

	var ebt = revenue - expenses

	var profit = revenue - expenses - ((revenue - expenses) * taxRate / 100)

	var ratio = ebt / profit

	EBT := fmt.Sprintf("ebt: %.2f\n", ebt)
	PROFIT := fmt.Sprintf("profit: %.2f\n", profit)
	RATIO := fmt.Sprintf("ratio: %.2f\n", ratio)

	fmt.Print(EBT, PROFIT, RATIO)
}

func outputText(text string) {

	fmt.Print(text)
}
