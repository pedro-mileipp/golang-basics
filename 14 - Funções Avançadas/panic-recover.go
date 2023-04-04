package main

import "fmt"
func recuperarExecucao(){
	if r := recover(); r != nil{
		fmt.Println("RECUPERADO")
	}
}


func alunoAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("MEDIA = 6") // a função panic vai interromper o fluxo do normal do programa e chamar todas as funções com defer, se não forem recuperadas as funções com o recovery o programa morre.

}

func main() {
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("Pós execução")
}