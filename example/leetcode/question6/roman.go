package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("MCMXCIV"))

}

func romanToInt(s string) int {

	output := 0
	roman := map[string]int{

		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	romanException := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	for j := 0; j < len(s); j++ {

		if j < len(s)-1 {

			exception, isOk := romanException[string(s[j])+string(s[j+1])]
			if isOk {
				s = string(s[:j]) + "**" + string(s[j+2:])
				fmt.Println(j, s, exception, isOk)
				output += exception
			}
		}

		value, ok := roman[string(s[j])]

		if ok {
			output += value
		}

	}

	//MCMXCIV

	// fmt.Println(s, output)
	// for _, word := range s {

	// 	value, ok := roman[string(word)]

	// 	if ok {
	// 		output += value
	// 	}

	// 	//fmt.Println(i, ok, value)
	// }

	//fmt.Println(output)

	return output

}
