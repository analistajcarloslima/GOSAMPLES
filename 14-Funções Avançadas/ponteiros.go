package main

import "fmt"

//Verificar como as funções podem usar ponteiros

func inverterSinal(numero int) int {
	return numero * -1
}

func invertercomponteiro(numero *int) { // agora ao inves de referenciar a variavel vou referenciar o endereço dela
	*numero = *numero * -1 //entao eu uso o valor que está dentro do endereço e multiplico por -1
}

func main() {
	numero := 20
	numeroinvertido := inverterSinal(numero)
	fmt.Println(numeroinvertido)
	fmt.Println(numero) // aqui podemos verificar que a variavel numero nao foi alterada
	// foi enviada uma copia dela para dentro da função inverter sinal
	novonumero := 40
	fmt.Println(novonumero)
	invertercomponteiro(&novonumero) // como eu vou enviar o endereço de memoria preciso suar o e comercial
	fmt.Println(novonumero)          // quando imprimir novamente vou ver que o valor da variavel foi alterado

}
