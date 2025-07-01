package main

import "fmt"

func fibonacci(n uint) uint {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	var anterior, atual uint = 0, 1

	for i := uint(2); i <= n; i++ {
		proximo := anterior + atual
		anterior = atual
		atual = proximo
	}

	return atual
}

func main() {
	fmt.Println("Fibonacci Iterativo")

	posicao := uint(10)

	for i := uint(0); i <= posicao; i++ {
		fmt.Printf("F(%d) = %d\n", i, fibonacci(i))
	}

}
