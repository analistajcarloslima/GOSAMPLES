package main //descreve que é um arquivo executavel
// todos os arquivos que estiverem declarados dentro do package main
//poderão se enchergar, mesma estrutura de pacotes java
//para o arquivo ser executavel obrigatoriamente ele precisa estar como package main
import ( //aqui em import voce colocar as extenções
	//pode colocar mais de uma extenção por linha dentro do import
	"fmt"
)

func main() { //para executar também é necessário uma função main
	fmt.Println("Olá mundo!")

}

//para executar o código basta abrir um terminar e digitar go run+ aruivo

//alteração
