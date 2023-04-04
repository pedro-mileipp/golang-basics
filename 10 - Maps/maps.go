package main

import "fmt"

func main() {
	fmt.Println("Maps")
	fmt.Println()

	usuario := map[string]string{ // dentro da chave o tipo da chave, fora o tipo dos valores
		"nome": "Pedro",
		"sobrenome": "Mileipp",
	}
	fmt.Println(usuario)
	
	// Imprimindos os valores através das chaves
	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])
	fmt.Println()


	usuario2 := map[string]map[string]string{ // map dentro de map
		"nome":{
			"primeiro": "Johnny",
			"ultimo": "Walker",
		},
		"curso":{
			"nome": "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2["nome"]["primeiro"]) // imprimindo o nome do usuario2
	fmt.Println()

	delete(usuario2, "nome") // deletando o campo nome
	fmt.Println(usuario2)
	fmt.Println()


	// adicionando campo no usuario2
	usuario2["materia"] = map[string]string{
		"nome": "Fundamentos de Arquitetura de Computadores",
	}
	fmt.Println(usuario2)



}