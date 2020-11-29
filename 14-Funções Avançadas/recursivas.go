package main

import "fmt"

func fibonacci(posicao uint) uint {
	// primeiro eu tenho que especificar as condições de parada
	//isso evita que a função seja executada muitas vezes causando o stackoverflow
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1) // no caso ele  ele vai receber o numero de entrada e somar com o numero -2 e o numero -1

}

func main() {
	fmt.Println("Recursivas")

	//Na matemática, a Sucessão de Fibonacci (ou Sequência de Fibonacci), é uma sequência de números inteiros, começando normalmente por 0 e 1, na qual, cada termo subsequente corresponde à soma dos dois anteriores. ... 0,1, 1, 2, 3, 5, 8, 13

	posicao := uint(6)
	fmt.Println(fibonacci(posicao))

	// não é recomendado usar quando temos muitas execuções a serem feitas. por exemplo 1000 repetições

	for i := uint(0); i <= posicao; i++ {
		fmt.Println(fibonacci(i))

	}

}
