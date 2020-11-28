package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	var numero int
	numero = 10

	if numero > 10 {
		fmt.Println("maior que 10")
	} else { // condição senão não é obrigatória
		fmt.Println("menor ou igual 10")

	}

	if var1 := numero; var1 > 0 { // posso fazer a declaração de variavel direto na condição
		fmt.Println("maior que 0")
	} //porem a variavel existe somente dentro da condição if não posso chamar ela fora
	// as vezes e necessario pois vc pode criar um variavel temporaria apenas para aquela condição

}
