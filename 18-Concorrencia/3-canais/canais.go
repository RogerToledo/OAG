package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Hello World!", canal)

	for mensagem := range canal {
		fmt.Println(mensagem)
	}
	
	fmt.Printf("Fim!!")

}

func escrever(texto string, canal chan string) {
	for i := 1; i <= 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}