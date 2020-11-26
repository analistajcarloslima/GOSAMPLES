package main

import "fmt"

func main() {
	//declarando variaveis
	var tipo1 string = "Variavel1" //primeira forma de declaração
	fmt.Println(tipo1)

	//Obs o go não deixa voce criar uma variável e não usar, dá erro de execução a mesma coisa vale para import
	tipo2 := "Variavel2" //segunda for de declarar a variavel de forma discreta, ele declara o tipo de variavel de acordo com o valor dela
	//declaração implicita (inferencia de tipo)
	fmt.Println(tipo2)

	// declaração multipla

	// com declaração

	var (
		tipo3 string = "Var3"
		tipo4 string = "Var4"
	)
	fmt.Println(tipo3, tipo4)

	//Declação multipla por inferencia
	tipo5, tipo6 := "Var5", "var6"
	fmt.Println(tipo5, tipo6)

	// declaração de constante
	const tipo7 string = "contante1"
	fmt.Println(tipo7)

	// no go e possivel trocar o valor das variaveis sem a necessidade de uma variavel auxilar para armazenar um valor temporario
	//inverter os valores de tipo5 e6
	tipo5, tipo6 = tipo6, tipo5
	fmt.Println(tipo5, tipo6)

}
