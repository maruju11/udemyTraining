package main

import "fmt"

func main() {
	fmt.Println("indice de massa corporal")
	z := indPeso(69.0, 1.77)
	fmt.Println(z)
	fmt.Println(indPeso(75, 1.77))
}

func indPeso(x, y float64) float64 {
	//var cpr float64
	cpr := x / (y * y)
	return cpr
	//fmt.Println(cpr)

}
