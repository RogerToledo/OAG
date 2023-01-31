package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wp sync.WaitGroup
	wp.Add(2)

	go func() {
		escrever("-> Goroutine")
		wp.Done()
	}()

	go func() {
		escrever("--> Goroutine")
		wp.Done()
	}()

	func () {
		escrever2("Testando")
	}()

	wp.Wait()
	
}

func escrever(texto string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(texto, i)
		time.Sleep(time.Second)
	}
}

func escrever2(texto string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(texto, i)
		time.Sleep(time.Second)
	}
}