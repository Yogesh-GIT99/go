package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strStr("sdsaddsad", "sad"))

}

func strStr(haystack string, needle string) int {

	if len(haystack) < len(needle) {
		return -1
	}
	k := len(needle) - 1
	if strings.Contains(haystack, needle) {
		for i := 0; i < len(haystack); i++ {
			if haystack[i:k+i] == needle {
				fmt.Println(haystack[i : k+i])
				return i
			}

		}
	}

	return -1
}
