package main

import (
	"fmt"
	"strings"
)

func main() {

	strs := []string{"aa"}
	fmt.Println(longestCommonPrefix(strs))

}

func longestCommonPrefix(strs []string) string {

	firstElement := strs[0]
	output := ""
	flag := false
	prefix := ""

	if len(strs) == 1 {
		return strs[0]
	}

	for _, word := range firstElement {
		prefix += string(word)
		for j := 1; j < len(strs); j++ {
			if strings.HasPrefix(string(strs[j]), prefix) {
				flag = true

			} else {
				return output
			}
		}
		if flag {
			output += string(word)
		}
	}

	return output
}
