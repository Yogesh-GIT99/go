package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {

	l := 0
	r := len(prices) - 1
	maxProfit := 0

	for l < r {

		profit := prices[r] - prices[l]

		// if prices[r] > prices[r+1] {
		// 	l++
		// 	continue
		// }

		if profit < 0 {
			l++
			continue
		}

		if profit > maxProfit {
			maxProfit = profit
		}

		r--
		fmt.Println("left: ", l, "right:", r)
	}
	return maxProfit
}
