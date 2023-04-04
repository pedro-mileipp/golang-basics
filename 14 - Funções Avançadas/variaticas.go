package main

import "fmt"

// São funções que podem receber uma quantidade n qualquer de parâmetros sem precisar serem especificadas

func soma(numeros ...int) int{
	total := 0
	for _, numero := range numeros{
		total += numero
	}
	return total
}

func main() {
	totalDaSoma := soma(10, 10, 10, 50)
	fmt.Print(totalDaSoma)
}


