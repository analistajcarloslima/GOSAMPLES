package main

import "fmt"

var result bool

// O panic paraq a execução do codigo independente do que vier depois caso uma consição for atingida

// uma maneira de recuperar a execução de um pacnic e usar o rescover
func recuperar() {
	if r := recover(); r != nil {
		fmt.Println("Função recuperada")
	}
}

func nota(n1, n2 float32) bool {
	defer recuperar() // antes do pacnic ser executado ele chamar todas as funções com defer
	// nesse momento eu uso a função recover que vai tenta rrecuperar a execução
	fmt.Println("Calculando nota")
	media := (n1 + n2) / 2

	if media > 6 {
		result = true // antes de declarar o retorno eu seto a variavel global
		return true
	} else if media < 6 { // o else tem que fica na mesma linha de fechamento do if
		result = false
		return false
	}

	panic("É EXETAMENTE 6") // basicamente se as condições não forem atendidadas ele vai parar o programa e dar uma msg de erro
}

func main() {

	media := nota(6, 6)
	fmt.Println(media)
	//caso a media seja ele não vai executar o restante do programa
	fmt.Println("fim do programa")

}
