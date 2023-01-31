package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Hello")
	escrever("World")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}