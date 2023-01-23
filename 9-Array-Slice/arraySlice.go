package main

import (
	"fmt"
)

func main() {
	fmt.Println("Array e Slice")

	var array1 [5]string
	array1[0] = "P1"
	fmt.Println(array1)

	array2 := [5]string{"Roger", "Ana", "Yasmin", "Kyara", "Wood"}
	fmt.Println(array2)

}