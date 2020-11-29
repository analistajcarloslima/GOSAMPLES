package main

import "fmt"

func main() {
	fmt.Println("funções anonimas")

	// basicamente são funções sem nome que podem ser auto executadas sem a ncessidade de chamada

	func(msg string) {
		fmt.Println(msg)

	}("passando parametro")

}
