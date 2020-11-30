package main

import "fmt"

// defer atrasa a execução de uma função
// defer quer dizer adiar

func funcao1() {
	fmt.Println("execuçao função1")

}
func funcao2() {
	fmt.Println("execuçao função2")

}

var result bool // aqui eu também aprendi como pegar o retorno de uma função e colocar numa variavel

//posso usar o defer para atrasar um comando dentro de uma função
func nota(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. O resultado é")
	fmt.Println("Calculando nota")
	media := (n1 + n2) / 2

	if media >= 6 {
		result = true // antes de declarar o retorno eu seto a variavel global
		return true
	}
	result = false
	return false

}

func main() {
	defer funcao1()
	fmt.Println("defer")

	//eu adio a execução da função para ultimo comando possivel
	funcao2()

	calc := nota(6, 5)
	fmt.Println(calc)
	fmt.Println(result)

}
