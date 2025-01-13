// os function / use of slices / len function / build a go binary
package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf("%#v\n", os.Args)

	fmt.Println("Path : ", os.Args[0])
	fmt.Println("1st argument : ", os.Args[1])
	fmt.Println("2nd argument : ", os.Args[2])

	fmt.Println("No. of items inside slice value : ", len(os.Args))

}

// go build -o build_name ( cmd to build a go binary)
