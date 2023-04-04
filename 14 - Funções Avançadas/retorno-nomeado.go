package main

import "fmt"

func CalculosMatematicos(n1, n2 int) (soma, subtracao int) { // retornara soma e subtração, já nomeadas na declaracao da função
	soma = n1 + n2
	subtracao = n1 - n2
	return
}
func main() {
	soma, subtracao := CalculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}