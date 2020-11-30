package main

import "fmt"

// Bascimaente o closure e uma função que retorna uma variavel que está em outra função

func closure() func() {
	texto := "Dentro da função closure"
	funcao := func() {
		fmt.Println(texto)

	}
	return funcao

}

func main() {
	texto := "Dentro da funcao main"
	fmt.Println(texto)
	novafuncao := closure()
	novafuncao() // basiment e pra comprovar que a função closure esta pegando a variavel interna dela
}
