package main

import (
	"errors"
	"fmt"
)

func main() {
	//tipos de dados
	//INTEIROS
	// temos tamanhos diferentes de bits pro  int int, int16, int32, int 64

	var numero int16 = 100
	fmt.Println(numero)
	//Obs se não especificar o tmaanho do int o Go pega o padrão da sua aquitetura x64 ou 32 bits
	//int também suporta  numeros negativos -

	//unsighnetd int
	//int que não suporta sinais positivos ou negativoas
	var numero2 uint16 = 100
	fmt.Println(numero2)

	//Alias
	//Alguns tipos de dados tem um apelido para facilitar a escrita
	// exemplo rune = int32
	//exemplo byte = uint8

	//números Reais //não suporta , tem que usar o .
	//pode ser float32 ou float 64
	var real1 float32 = 123.45
	fmt.Println(real1)

	//Texto String
	var texto string = "string"
	fmt.Println(texto)

	//go não tem Char o que seria o char usando aspas simples ele transforma na representação do caracter na tabela ASCII
	//exemplo
	char := 'B'
	fmt.Println(char)

	//Valor de inicialização de variavel
	// No go não é obrigatorio inicializar a variavel
	//ele atribui valores padrão de acordo com o tipo de variavel
	// Para texto ele atribui um char vazio
	//para numero ele atribui 0
	//Exemplo

	var ini int32
	fmt.Println(ini)

	//Booleano

	var boooleano bool //o padrão de inicialização e false
	fmt.Println(boooleano)

	//No go error é um tipo de dado para tratamento especial
	//O Padrão de inicialização e nill= nulo que equivale a zero
	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro interno") //posso enviar um parametro para dentro da variavel de erro
	fmt.Println(erro2)
}
