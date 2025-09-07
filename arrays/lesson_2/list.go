package main

import "fmt"

func main() {
	platforms := [4]string{"patna", "jhodpur", "srinagar"}
	fmt.Println(platforms[1:2]) // slicing 1st and last
	// other ways
	fmt.Println(platforms[:2])
	fmt.Println(platforms[1:]) // we cannot use -1 here unlike other programs and cannot use index value +1 more than the last index

	// build another slice from slice
	station := platforms[1:]
	stations := station[:1]

	fmt.Println(stations)
}
