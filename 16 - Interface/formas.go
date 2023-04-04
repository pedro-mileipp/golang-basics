package main

import (
	"fmt"
	"math"
)

// Usando a interface não vai ser preciso escrever mais de um metodo para escrever a area
type forma interface {
	area() float64
}

// Para usar esse método cada struct (nesse contexto, retangulo e circulo) deverão ser formas, para serem formas, precisam ter um método chamado area que não recebe parâmetros e retorna um float64
func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)

}