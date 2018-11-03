package main

import "fmt"
impr(
	L "./media"
)

func main() {
	fmt.Println("indice de massa corporal")
	z := indPeso(69.0, 1.77)
	fmt.Println(z)
	fmt.Println(indPeso(75, 1.77))
	fmt.Println("--------------------------------------------------------------------------")
	z2 := []float64{100, 0200, 150, 300, 220, 310, 222}
	fmt.Println(L.Media(z2))
}

func indPeso(x, y float64) float64 {
	//var cpr float64
	cpr := x / (y * y)
	return cpr
	//fmt.Println(cpr)

}
