package main

import(
	"fmt"
)

func main() {
	/* No Go há duas maneiras de declarar variáveis, com o tipo explícito e implícito */
	
	// Explícito
	var s1 string = "Essa variável foi declarada com o tipo explícito"
	fmt.Println(s1);

	// Implícito 
	s2 := "Essa variável foi declarada com o tipo implícito" // Inferência de tipo, inferindo o tipo de variável com base no seu valor atribuído.
	fmt.Println(s2);
	fmt.Println()
	

	// Declarando múltiplas variáveis
	var(
		mensagem1 = "Primeira forma de declarar múltiplas variáveis:"
		n1 int = 1
		n2 int = 2
	)
	fmt.Println(mensagem1,n1, n2)

	mensagem2, n3, n4 := "Segunda forma de declarar múltiplas variáveis:", 3, 4
	fmt.Println(mensagem2,n3, n4)
	fmt.Println()

	const c1 int = 1
	fmt.Printf("A constante c1 de valor %d não pode ser alterada", c1) // O Parintf suporta melhor formatação
	
}