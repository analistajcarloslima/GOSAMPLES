package main

import "fmt"

func main() {
	fmt.Println("ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1
	//A PARTIR DAQUI SE EU ALTERAR O VALOR DE VARIAVEL 1 A VAR2 NÃO MUDA POIS ELA COPIOU O
	//CONTEUDO DE V1

	fmt.Println(variavel1, variavel2)

	//A IADEIA E QUE EU POSSA EM VEZ DE REFERENCIAR O CONTEUDO EU REFERENCIE O
	//ENDEREÇO DE MEMORIA DE V1

	//PONTEIRO E UMA REFERENCIA DE MEMORIA
	var var3 int
	var ponteiro *int

	var3 = 100
	ponteiro = &var3 // aqui eu faço o ponteiro apontar pro endereço da var3
	fmt.Println(var3, ponteiro)
	//caso eu quei ver o conteudo que esta dentro do ponteiro eu uso o *
	//desreferenciação
	fmt.Println(*ponteiro)

}
