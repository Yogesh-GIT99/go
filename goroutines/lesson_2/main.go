package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChan := make([]chan bool, len(taxRates)) // declaring channel
	errorChan := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		doneChan[index] = make(chan bool) // intializing channel
		errorChan[index] = make(chan error)
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChan[index], errorChan[index])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	for index := range taxRates { // using select to handle errChan channel
		select {
		case err := <-errorChan[index]:
			if err != nil {
				fmt.Print(err)
			}
		case <-doneChan[index]:
			fmt.Println("done !")
		}
	}
}
