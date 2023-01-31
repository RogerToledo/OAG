package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	canal := escrever("Hello World!!")
	fmt.Println(<-canal)

	for i := 0; i < 10; i++ {
		canal := escrever(strconv.Itoa(i))
		fmt.Println(<-canal)
	}

}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido, %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}