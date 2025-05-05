package main

import (
	"fmt"
	"os"
)

func main() {
	revenue := getUserInput("Revenue: ")
	if revenue == 0 {
		fmt.Println("You have entered wrong value for revenue multiple times! bye")
		return
	}
	expenses := getUserInput("Expenses: ")
	if expenses == 0 {
		fmt.Println("You have entered wrong value for expenses multiple times! bye")
		return
	}
	taxRate := getUserInput("Tax Rate: ")
	if taxRate == 0 {
		fmt.Println("You have entered wrong value for taxRate multiple times! bye")
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

	writeResults(ebt, profit, ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		for i := 0; i < 1; i++ {

			fmt.Println("Invalid value! please enter a value greater than zero.")
			fmt.Print(infoText)
			fmt.Scan(&userInput)
			if userInput > 0 {
				return userInput
			}

		}
		return 0
	}
	return userInput
}

func writeResults(ebt, profit, ratio float64) {

	// precise := math.Pow(10, float64(2))
	// ratio1 := math.Round(ratio*precise) / precise
	text := fmt.Sprintf("ebt: %.1f\nprofit: %.1f\nratio: %.1f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(text), 0644)

}
