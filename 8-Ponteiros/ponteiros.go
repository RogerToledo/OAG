package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Printf("Variável 1: %d\n", variavel1)
	fmt.Printf("Variável 2: %d\n", variavel2)

	variavel1 = 20

	fmt.Println("Variável 1 alterada para 20")
	fmt.Printf("Variável 1: %d\n", variavel1)
	fmt.Printf("Variável 2: %d\n", variavel2)

	fmt.Println("------------------------------------------------")

	var variavel3 int = 100
	var variavel4 *int = &variavel3

	fmt.Printf("Variável 3: %d\n", variavel3)
	fmt.Printf("Variável 4: %d\n", *variavel4)
	
	variavel3 = 200

	fmt.Println("Variável 3 alterada para 200")
	fmt.Printf("Variável 3: %d\n", variavel3)
	fmt.Printf("Variável 4: %d\n", *variavel4)

	*variavel4 = 300
	fmt.Println("Variável 4 alterada para 300")
	fmt.Printf("Variável 3: %d\n", variavel3)
	fmt.Printf("Variável 4: %d\n", *variavel4)

}