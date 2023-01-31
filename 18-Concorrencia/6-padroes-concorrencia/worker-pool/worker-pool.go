package main

import "fmt"

func main(){
	tarefas := make(chan int, 45)
	resultado := make(chan int, 45)

	go worker(tarefas, resultado)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultado
		fmt.Println(resultado)
	}
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for numeros := range tarefas {
		resultados <- fibonacci(numeros)
	}

}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao - 2) + fibonacci(posicao - 1)
}