package app

import (
	"github.com/urfave/cli"
)

//Gerar -defiinir o nome da função com maisculo para ela poder ser exportada para o package main, ele vai criar uma aplicação de linha de comando
func Gerar() *cli.App {
	app := cli.newApp()
	app.Name = "Aplicação linha de comando"
	app.Usage = "Busca de IPs e Nomes de servidor na internet"
	return app
}
