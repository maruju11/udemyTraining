package main

import (
	"fmt"
	"math"
)

// Square = um tipo definido pelo usuario
type Square struct {
	lado float64
}

//Circle polimorfismo
type Circle struct {
	raio float64
}

func (l Square) area() float64 {
	r := l.lado * l.lado
	return r
}
func (c Circle) area() float64 {
	return math.Pi * c.raio * c.raio
}

// Shape possui a mesma assinatura do método area(r float64)
type Shape interface {
	area() float64
}

func info(s Shape) {
	fmt.Println(s)
	fmt.Println(s.area())
}

func main() {
	ss := Square{4}
	sc := Circle{4}
	info(ss)
	info(sc)
}
