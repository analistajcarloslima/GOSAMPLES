package main

import "fmt"

//Funções também são tipos de dados
//elas podem ter parametros e retornos
//parametros entrada retorno= saida

func somar(n1 int8, n2 int8) int8 { //entre parenteses coloca os parametros //fora colocar o tipo de retorno
	return n1 + n2
}

func main() {
	soma := somar(10, 20) //soma vai receber o retorno da função somar
	fmt.Println(soma)

	//também posso atribuir uma função para uma variavel
	//para que ela seja executada com os valores definidos na entrada
	var f = func(txt string) string {
		fmt.Println(txt)
		return "retorno"
	}
	//crio uma variavel que recebe a função f passando os parametros de entrada
	resultado := f("Parametro de entrada")
	fmt.Println(resultado)

	// para chamar uma função com mais de um retorno preciso declarar um variavel para cada retorno
	resultadosoma, resultadosub := calculos(10, 15)
	fmt.Println(resultadosoma, resultadosub)

	//Caso eu não queira todos os retornos de uma função posso usar o o _ na chamada
	//Exemplo:
	somentesoma, _ := calculos(10, 15)
	fmt.Println(somentesoma)

}

//Obs no Go as funções podem ter mais de um retorno

func calculos(n1, n2 int8) (int8, int8) { //aqui eu especifico a quantidade e is tipos de retorno
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub //a ordem dos retornos define os resultados devolvidos pela função

}
