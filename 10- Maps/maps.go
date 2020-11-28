package main

import "fmt"

func main() {
	fmt.Println("maps")

	// maps e bem pareceido com strutct classes, mas é mais rigido com relaçao ao tipo de dados
	// criando um map usuario
	// as chaves e os valore sprecisam ser do mesmo tipo
	usuario := map[string]string{ // funciona parecido com BD vc tem um tipo de chave e um tipo de valor
		"nome":      "pedro", // aqui nome vai ser a chave e pedro o valor
		"sobrenome": "silva",
	}

	fmt.Println(usuario)
	// para chamar os campos especificos eu uso o nome da chave
	fmt.Println(usuario["sobrenome"])

}
