package main

import (
	"fmt"
	"modulo/auxiliar" // é preciso fazer a importação do pacote auxilar, chamando o módulo primeiro

	"github.com/badoux/checkmail" //fazendo a importação de um módulo externo
)

func main() {
	fmt.Println("Escrevendo arquivo no main")
	auxiliar.Escrever() // aqui eu fiz a chmada da função Escrever que está no arquivo tauxilar
	//não preciso chamar examente o nome do arquivo apenas do pacote
	//escrever está trazendo a função publica Escrever e também a função privada escrever2 do arquivo escrever2.go

	erro := checkmail.ValidateFormat("123") //verfica formato valido de email
	fmt.Println(erro)
}

// fara fazer download de pacotes externos use o terminal dentro da pasta raiz do projeto:
//usando o comando go get
//exemplo go get github.com/badoux/checkmail
