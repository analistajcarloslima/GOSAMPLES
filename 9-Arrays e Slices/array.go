package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	// arrays são estruturas de dados
	//no Go e necessario especificar o tamanho e o tipo de dados
	// não pode haver dados diferente do epecificado

	// array de int
	var array1 [5]int8
	fmt.Println(array1)

	//array de string
	var array2 [5]string
	// podemos popular diretamente pela posição do array
	array2[0] = "posição1"
	fmt.Println(array2)

	//Também posso criar um array por inferencia
	array3 := [3]string{"1", "2", "3"}
	fmt.Println(array3)

	//Tambem posso definir o tamanho do arary pela quantidade de valores
	array4 := [...]int{1, 2, 3, 4} //o ... define o tamanho do array pela quantidade de entradas
	fmt.Println(array4)
	fmt.Println(array4[3])

	//SLICES SÇAO ESTRUTURAS MAIS FLEXIVEIS COM RELAÇÃO A TAMANHO
	slice1 := []int{1, 2, 3, 4} // basicamente a mesma coisa do array são não preciso especificar o tmaanho
	fmt.Println(slice1)
	// para provar que o slice e o array são deferentes podemos usar uma função que retorna o tipo de uma variavel : reflect
	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array4))

	//para adicinar um novo valor ao slice usa o append
	slice1 = append(slice1, 5) // ele recebe o valor atual do slice e retona um novo slice atualizado
	//poderia criar um novo slice nessa chamada do append
	fmt.Println(slice1)

	// como o slice usa a estrutura do aray eu posso pegar parte de um array e colocar dentro de um slice
	slice2 := array2[0] // asicamente ele aponta para posições do array
	fmt.Println(slice2)

	//Arrays internos
	// bascamente quando um clice é criado ele referencia um array interno que não é mostrado
	//podemos ver a criação desse array através da função make
	fmt.Println("-------Array interno-------")

	slice3 := make([]int, 10, 11) // basicamente o make criar um array de tmaaho 15 e pega uma fatia das primeiras 10 posições e coloca no slice3
	fmt.Println(slice3)
	//para eu ver o tamanho geral do slice eu uso o lenght
	fmt.Println(len(slice3))
	//para eu ver a capacidade geral do slice eu uso o cap
	fmt.Println(cap(slice3))
	// caso a capacidade maxima do slice seja estourada o go cria um novo array com o dobro de tamanho para referenciar
	slice3 = append(slice3, 12)
	fmt.Println(slice3)
	slice3 = append(slice3, 13)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3)) // ele dobra a capacidade do array refenciado para 22

	slice4 := make([]int, 3) // caso eu não queira definir a capacidade maxima ele define a capacidade pelas posições
	fmt.Println(slice4)
	slice4 = append(slice4, 13) // aqui eu extrapolo  a quantidade para ele dobrar a quantidade do array
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
