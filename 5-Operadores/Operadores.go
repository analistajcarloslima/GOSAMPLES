package main

import "fmt"

func main() {
	//OPERADORES ARITIMÉTICOS
	var soma int32 = 1 + 2
	multi := 1.5 * 3
	sub := 1 - 2
	div := 1 / 2
	mod := 2 % 2

	fmt.Println(soma, multi, sub, div, mod)

	// Obs no go não se pode fazer operações com tipos de dados diferentes
	//exemplo não posso somar um int32 com um int64
	var numero1 int16
	var numero2 int16

	var soma2 = numero1 + numero2
	fmt.Println(soma2)

	//OPERADORES DE ATRIBUIÇÃO
	var variavel1 string = "string" //por tipo
	variavel2 := "string2"          //por inferencia
	fmt.Println(variavel1, variavel2)

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2) //comparação
	fmt.Println(1 != 2) // diferente

	// OPERADORES LÓGICOS-TABELA VERDADE
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) //AND
	fmt.Println(verdadeiro || falso) //OR
	fmt.Println(!verdadeiro)         //NOT

	//OPERADORES UNÁRIOS
	numero3 := 10
	numero3++ //incremento de 1
	fmt.Println(numero3)
	numero3 += 15 //numero = numero +15 incrementa ele mesmo mais um numero
	fmt.Println(numero3)

	numero3--     //decremento
	numero3 -= 15 //umero = numero -15 decrementa ele mesmo mais um numero
	fmt.Println(numero3)

	//OPERADOR TERNÁRIO
	//abribuir o valor de uma variavel baseado em uma condiição na declaração
	// no go isso não é possível
	//exemplo:
	// texto := numero >8 ? "maior que 5" : "menor que 5"

	//  para faze risso e necessario implementar uma condição

	var texto string
	if numero3 > 5 {
		texto = "maior que 5"
	} else {
		texto = "menor que 5"
	}
	fmt.Println(texto)

}
