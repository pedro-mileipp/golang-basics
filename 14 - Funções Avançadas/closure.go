package main

import "fmt"

// São funções que referenciam variáveis fora do seu corpo


func testClosure() func(){
	texto := "Texto dentro da closure"

	funcao := func(){
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da func main"
	fmt.Println(texto)
	fmt.Println()

	funcNova := testClosure()
	funcNova()
}