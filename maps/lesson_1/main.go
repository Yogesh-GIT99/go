package main

import "fmt"

func main() {

	games := map[string]string{
		"war of raganrok": "2020",
		"fifa mix":        "2025"}

	fmt.Println(games)

	// get a specific value from map
	fmt.Println(games["fifa mix"])

	// adding/overwriting value to map
	games["wwe"] = "2026"
	fmt.Println(games)

	// delete a value from map
	delete(games, "fifa mix")
	fmt.Println(games)
}
