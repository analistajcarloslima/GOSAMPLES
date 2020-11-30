package main

import "fmt"

// basicamente o função init é executada antes do main
// pode ser usado em arquivos do mesmo pacote pra inicar a execução de um arquivo

func main() {
	fmt.Println("função main")

}

// não importa a ordem
func init() {
	fmt.Println("função init")
}
