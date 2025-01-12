// learn how to use path seperator function and define variables

package main

import (
	"fmt"
	"path"
)

func main() {

	var dir, file string

	dir, file = path.Split("api/main/index.html")

	fmt.Println("dir : ", dir)
	fmt.Println("file : ", file)
}

// func main() {

// 	var file string

// 	_, file = path.Split("api/main/index.html") // can also get the value of one variable like this

// 	fmt.Println("file : ", file)
// }

// func main() {

// 	dir, file := path.Split("api/main/index.html") // short way of declaring variables

// 	fmt.Println("file : ", file)
// 	fmt.Println("dir : ", dir)
// }

// Note: Use short declaration when you know the value before hand and using variables is handy when you want to make some code readable and do not know
//       the value before hand.
