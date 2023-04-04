package main

import(
	"fmt"
)

func somar(n1 int8, n2 int8) int8{ // declarando funcão
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) { // função com múltiplos retornos
	var soma int8 = n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(50, 50) // chamada da função
	fmt.Println(soma)
	
	var f = func(txt string){ // variável que recebe uma função
		fmt.Println(txt)
	}

	f("Texto da função f")
	
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}