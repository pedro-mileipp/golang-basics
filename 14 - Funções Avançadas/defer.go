package main

import(	
	"fmt"
)

// O defer adia a execução de determinado pedaço de código


func funcao1(){
	fmt.Println("Executando a func1")
}

func funcao2(){
	fmt.Println("Executando a func2")
}

func funcao3(){
	fmt.Println("Executando a func3")
}

func main() {
	defer funcao2() // ultimo, primeiro defer sempre será o últikmo
	defer funcao3() // segundo, depois do primeiro defer os outros vem posteriormente
	funcao1() // primeiro 
}