package main

import "fmt"

// structs funcionam como as classes de Objetos

//criando uma struct

type usuario struct {
	nome     string
	idade    uint8
	endereço endereço // aqui fiz a importação de uma struct pra dentro do struct usuario
}

func main() {
	fmt.Println("aruivo struct")

	//pra instanciar uma struct eu declaro uma variavel do tipo da struct
	var u usuario
	//para inserir dados eu uso os atributos da variavel
	u.nome = "João"
	u.idade = 10
	fmt.Println(u.nome) // na saida posso chamar os atributos de forma separada

	//também possso estanciar a struct por inferencia
	endereçoexemplo := endereço{"rua01", 12}
	u2 := usuario{"Jose", 10, endereçoexemplo}
	fmt.Println(u2)

	//também possso estanciar sem passar todos os valores
	u3 := usuario{idade: 10}
	fmt.Println(u3)

}

//Eu posso ter structs dentro de structs
//exemplo vouc riar uma struct para endereço e atribuir a stuct usuario
type endereço struct {
	rua    string
	numero uint8
}
