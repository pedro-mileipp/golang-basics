package main

import "fmt"

type endereco struct {
	logradouro string
	numero     uint8
}

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}


func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	u.endereco = endereco{"Rua X,", 35}
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos,", 0}

	usuario2 := usuario{"Leo", 21, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 22, nome: "Miguel", endereco: endereco{"Rua Y,", 32}}
	fmt.Println(usuario3)
}