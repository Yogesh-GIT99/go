package getPrices

import (
	"encoding/json"
	taxCalc "example/priceCalculator/calculation"
	"fmt"
	"os"
)

type data struct {
	Price []float64 `json:"price"`
}

func GetPrices(file string) []float64 {

	priceData, err := os.ReadFile(file)

	if err != nil {
		panic(err)
	}

	var resultPrice data
	err = json.Unmarshal(priceData, &resultPrice)

	if err != nil {
		panic(err)
	}

	return resultPrice.Price
}

func WriteTaxedfile(taxedPrice taxCalc.Pricematrix) {

	file, err := os.Create("output.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	for _, row := range taxedPrice {
		for _, val := range row {
			fmt.Fprintf(file, "%.f  ", val)
		}
		fmt.Fprintln(file)

	}

	fmt.Println(taxedPrice)

}
