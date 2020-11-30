package main

import (
	"fmt"
	"math"
)

//Basicamente as interfaces servem para criar métodos mais genéricos

//criação da interface é necesspario definir um método que toda função que quiser implementar a interface vai ter que ter
//o conteudo desse método pode variar assim eu posso ter uma maior flexibilidade

type forma interface {
	area() float64
}

// aqui eu defino o método que vai receber a interface como entrada
func escreverArea(f forma) {
	fmt.Printf("a área da forma é %0.2f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 { // cada método precisa implementar a interface com o metodo area()
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 { // cada método precisa implementar a interface com o metodo area()
	return math.Pi * math.Pow(c.raio, 2) //biblioteca math tem varias funções matematicas
	//math. Pi valor da contante pi
	//math.Pow faz a exponenciação
}

func main() {

	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
	// nesse caso eu posso usar o método escrever área para ambos os structs
	//através da interface
}
