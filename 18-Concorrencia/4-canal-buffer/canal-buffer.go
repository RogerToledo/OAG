package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Olá mundo"
	canal <- "hello world"
	canal <- "deadloack"

	mensagem  := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)

}