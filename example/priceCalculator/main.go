package main

import (
	taxCalc "example/priceCalculator/calculation"
	fmanager "example/priceCalculator/fileManager"
)

func main() {

	file := "price.json"

	prices := fmanager.GetPrices(file)

	taxedPrice := taxCalc.TaxCalc(prices)

	fmanager.WriteTaxedfile(taxedPrice)

}
