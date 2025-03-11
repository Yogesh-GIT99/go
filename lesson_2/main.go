// Profit calculator

package main

import "fmt"

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("taxRate: ")
	fmt.Scan(&taxRate)

	var ebt = revenue - expenses

	var profit = revenue - expenses - ((revenue - expenses) * taxRate / 100)

	// var profit = ebt * ( 1 - (taxRate/100))

	var ratio = ebt / profit

	// fmt.Println("ebt:", ebt, " profit:", profit, " ratio:", ratio)
	fmt.Printf("ebt: %.2f\nprofit: %.2f\nratio: %.2f\n", ebt, profit, ratio)
}
