package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}
type estudante struct {
	pessoa        // para fazer a herança basta declarar a strut a qual quer herdar
	curso  string // adiciona somente os atribuidos especificos da classe
}

func main() {
	fmt.Println("Herança")
	p1 := pessoa{"João", "Pedro", 20, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "engenharia"} // na instancia eu posso preeencher com os dados de uma outra classe
	//no caso e1 herdou os valores de p1 e adicionou os atributos de curso
	fmt.Println(e1)

}
