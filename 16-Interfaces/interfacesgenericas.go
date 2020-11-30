package main

import "fmt"

//interfaces genericas serem pra receber qualquer tipo de dados

//vou criar uma função que vai receber uma interface genérica

func generica(interf interface{}) {
	fmt.Println(interf) //criei um método que rece uma interface que não faz nada e retorna ela mesma
}

func main() {
	generica("string")
	generica(1)
	generica(true)
	//print ln também implementa uma interface generica que pode receber qualquer coisa
	fmt.Println("string", 1, true)
}
