package main

import "fmt"

// go não é orientado a objetos mas ele fornece suporte a metodos atraves das funcoes

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() { // esse u e apenas uma variavel que vai respresentar o objeto que está chamadndo o metodo no caso usuario
	// usando o Printf posso chamar tipos de dados diretamente na exibição
	fmt.Printf("usuario %s salvo com sucesso \n", u.nome) // no caso eu  uso o %s tambem preciso usar o \n para quebrar a linha
}

// vouc riar uma função para verificar a idade do usuario
func (u usuario) verificaidade() bool {
	return u.idade >= 18

}

// vou criar um função que vai adicionar um ano a idade do usuario
// toda vez que for necessario alterar os dados da struct usa o * ponteiro para refenciar
func (u *usuario) addicionaidade() { //atenção aqui pois fiz refencia ao ponteiro de usuario
	u.idade++
}

func main() {
	usuario1 := usuario{"jose", 20}
	fmt.Println(usuario1)

	usuario1.salvar() // aqui eu faço a cahamada da função salvar como se fosse um método
	idade := usuario1.verificaidade()
	fmt.Println(idade)

	usuario2 := usuario{"joao", 30}
	usuario2.addicionaidade()
	fmt.Println(usuario2.idade)

}
