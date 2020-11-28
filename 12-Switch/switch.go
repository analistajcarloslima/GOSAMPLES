package main

import "fmt"

//swittch é usado quando temos muitas condições

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(5) // aqui eu crio uma varivaél que vai cahmar a função dia da seman
	fmt.Println(dia)

	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)

	dia3 := diaDaSemana3(2)
	fmt.Println(dia3)

}

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"

	case 2:
		return "Segunda"
	// caso nenhuma das clausulas for atendida ele vai retornar o default
	default:
		return "numero inválido"
	}

}

// outra forma de fazer a validação e a comparação pelo case
func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda"
	default:
		return "número inválido"
	}
}

// outra forma de fazer e jogando o resultado da vailidação em uma variavel
func diaDaSemana3(numero int) string {
	var semana string
	switch {
	case numero == 1:
		semana = "domingo"
	case numero == 2:
		semana = "segunda"
	default:
		return "número inválido"
	}
	return semana // no caso ele vai retornar o valor da viavel semana
}
