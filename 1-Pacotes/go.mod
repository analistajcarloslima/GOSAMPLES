module modulo

go 1.15

// os módulos são estruturas que ornizam todas as depencias de um pacote
// quando geramos o executavel do módulo
//ele pega os binário de todos os aruivos e transforma em um único
//executavel
//ele também centraliza todas as dependias do projeto
//para criar um módulo execute esse comando no terminal:
// go mod init modulo

// obs tem que criar pelo terminal dentro da pasta do pacote

//para compilar o código em um unico arquivo atraves do módulo
//execute o comando
//go build
//sera gerado um arquivo com o nome do módulo
// para executar o arquivo gerado  use  .\modulo.exe dentro do diretório do arquivo

// o arquivo executavel não atualiza automaticamente, é necesário fazer o build novamente

require github.com/badoux/checkmail v1.2.1 // indirect
