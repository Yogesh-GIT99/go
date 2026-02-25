package taxCalc

import "fmt"

type Pricematrix [][]float64

func TaxCalc(price []float64) Pricematrix {
	taxRates := []float64{0, 10, 20}
	taxedPrice := make(Pricematrix, len(taxRates)) // creating outer slice rows [[] [] []]
	fmt.Println("matrix: ", taxedPrice)

	for i, rate := range taxRates {
		row := make([]float64, len(price)) // creating inner slice rows [0 0 0]
		fmt.Println("row: ", row)
		for j, p := range price {
			row[j] = p + (p * rate / 100)
		}
		taxedPrice[i] = row // inserting commplete row into matrix [[10 20 30] [] []]
		fmt.Println("taxedPrice: ", taxedPrice)
	}

	return taxedPrice
}
