package main

import (
	"fmt"
	"time"
)

// obs no go só tem a estrutura for

func main() {
	i := 0

	for i < 5 { //enquanto a condição for verdadeira ele vai executar
		i++
		fmt.Println("incrementando i")
		time.Sleep(time.Second) // time sleep pausa o codigo em um segundo
		fmt.Println(i)

	}
	fmt.Println(i)

	// outra forma de declarar
	for j := 0; j < 10; j += 2 { //tambem posso incrementar em unidades definidas com +=
		fmt.Println("incrementando j", j)
		time.Sleep(time.Second) // time sleep pausa o codigo em um segundo
		fmt.Println(j)
	}
	// caso eu deckare por inferencia não posso usar a variavel de repetição
	//fmt.Println(j) so vai poder ser usada dentro do for

	//For com Range -- usado para percorrer estruturas
	nomes := [3]string{"joão", "davi", "lucas"}

	for indice, nome := range nomes { //indice é a posição no array e nome e o valor que vai retornar
		fmt.Println(indice, nome)
	}

	//caso eu não quisesse mostrar a posição do indice basta usar o _
	for _, nome2 := range nomes { //indice é a posição no array e nome e o valor que vai retornar
		fmt.Println(nome2)
	}

	//Também posso usar o for para iterar com uma string

	for indice2, letra := range "PALAVRA" {
		fmt.Println(indice2, string(letra)) //caso eu não colocasse a clausula string ia retornar o Valor ASCII
	}

	//Também posso usar o for para iterar sobre maps
	usuario := map[string]string{
		"nome":      "leonardo",
		"sobrenome": "silva",
	}
	fmt.Println(usuario)

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
	//Não posso usar o range em structs

	// PAra criar um loop infinito basta eu não especificar condições
	//for {
	//	fmt.Println("Executanto inficitamente")
	//	time.Sleep(time.second)
	//}
}
