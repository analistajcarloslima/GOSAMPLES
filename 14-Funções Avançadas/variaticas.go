package main

import "fmt"

// funções variaticas são flexiveis com relação a quantidade de valores de entrada

func soma(numeros ...int) int { // usando o ... eu não preciso especificar a quantidade de valores de entrada
	fmt.Println(numeros) // ele cria um array com os numeros
	total := 0
	// agora vou percocer o arry com o for para fazer a soma dos valores
	for _, numero := range numeros {
		//vai rodar até passsar por todos os numeros de entrada
		total += numero //total vai ser igual total+ numero atual
	}
	return total
}

// outro exemplo de função
// essa recebe um valor fixo de string que sempre vai ser o mesmo
// e recebe um valor variavel de números
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}

}

// função de boas vindas

func boasvindas(texto string, nomes ...string) {
	for _, boas := range nomes {
		fmt.Println(texto, boas)
	}

}

func main() {
	totalsoma := soma(10, 20, 30) // façoa a chamada da função com a quantidade variavel de valores
	fmt.Println(totalsoma)

	//Observações no go só pode haver um parametro variavel com ...
	// e ele tem que ser o ultimo parametro especificado.
	// func soma(numeros ...int)

	escrever("olá mundo", 10, 20, 30) // como a função não tem retorno pósso chamar sem a variavel

	boasvindas("olá ", "joão", "jose", "pedro")
}
