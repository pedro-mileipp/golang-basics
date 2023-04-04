package main

import(
	"modulo/auxiliar"
	"fmt"
	"github.com/badoux/checkmail" // pacote externo
)

// O módulo é um conjunto de pacotes que compõem o projeto, sejam os pacotes criados pelo desenvolvedor, pacotes externos e pacotes do Go.
func main() {
	fmt.Println("Olá mundo!")
	auxiliar.Escrever() // função que está na pasta auxiliar é chamado no arquivo principal
	
	auxiliar.PularLinha()
	fmt.Println("OBS: As letras de funções em módulos precisam ser maiúsculas para serem vísiveis para arquivos de outro pacote")


	// usando pacote externo
	test := checkmail.ValidateFormat("pedromileipp@gmail.com") 
	fmt.Println(test)
}	