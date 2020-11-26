package auxiliar // nesse caso estou direcinando esse aquivo pra pasta auxiliar que criei na interface
//não importa o nome do aruivo o que importa e pretar atenção na estrutura de pastas

import "fmt"

// o nome da funções deve começar em maísculo para poder ser enchergada de outros pacotes
//como o go não é orientado a objetos ele usa a primeira letra da função para definir se vai ser public ou private

//Escrever é uma função que escreve na tela
func Escrever() {
	fmt.Println("Escrevendo do auxiliar2")
	escrever2() //como a função privada escrever 2 está no mesmo pacoteauxilar, posso chamar ele direto

}

//toda função que é exportada precisa ter um comentário acima descrevendo o que essa função faz
