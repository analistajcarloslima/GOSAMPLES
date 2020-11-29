package main

import "fmt"

// no retorno nomedado eu posso declarar o nome das variaveis de retorno
func calculosMatematicos(n1, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return // no caso ele vai retornar as duas variaveis nomeadas
}

func main() {
	// para chamar essas função eu preciso de duas variaveis para armazenar os valores
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)

}
