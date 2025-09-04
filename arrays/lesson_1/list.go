package main

import "fmt"

func main() {
	platforms := [4]string{"patna", "jhodpur", "srinagar"}
	var tickets [2]string = [2]string{"plat1"}
	fmt.Println(platforms)

	platforms[2] = "punjab"

	fmt.Println(platforms[2])
	fmt.Println(tickets)
}
