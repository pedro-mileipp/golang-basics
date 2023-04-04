package main

import (
	"fmt"
	"sync"
	"time"
)


func main() {
	
	var waitGroup sync.WaitGroup
	
	waitGroup.Add(2)
	
	
	go escrever("Olá Mundo!" ) // goroutlne
	escrever("Programando em Go!" )
}


func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time. Second)
	}
}


/* CONCORRÊNCIA ! = PARALELISMO
Concorrência é basicamente a capacidade de lidar com várias coisas de uma só vez, enquanto paralelismo é a capacidade de lidar com várias coisas ao mesmo tempo.

Imagine que você está programando. Ao finalizar aquela função lindona, você para, pega sua jarra de café, dá um gole, volta a jarra para a mesa e volta a programar na sequência. Nesse cenário você está programando e tomando café, ou seja, está fazendo várias coisas. Isso é concorrência.

Agora imagine que você está de fone ouvindo aquele som que te faz sentir o próprio Mr. Robot. Enquanto ouve esse som, você digita freneticamente em seu teclado mecânico cheio de luzes, programando uma média de 200 linhas por minuto… Nesse cenário, você está programando e ouvindo música ao mesmo tempo. Isso é paralelismo. */