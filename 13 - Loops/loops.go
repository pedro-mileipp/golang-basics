package main

import (
	"fmt"
	// "time"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Printf("Incrementando i, i = %d\n", i)
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println(i)

	// for j := 0; j < 10; j += 5 {
	// 	fmt.Printf("Incrementando j, j = %d\n", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"João", "Davi", "Lucas"}

	for k, nome := range nomes {
		fmt.Printf("%d - %s\n", k, nome)
	}

	for _, letra := range "PALAVRA" {
		fmt.Println(string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// for {
	// 	fmt.Println("Executando infinitamente")
	// 	time.Sleep(time.Second)
	// }
}