package main

import "fmt"

func main() {
	var a float64 = 0
	var p float64
	fmt.Print("digite a altura: ")
	fmt.Scanf("%f", &a)
	//fmt.Println("2a parte")
	fmt.Print("digite o peso ")
	fmt.Scanf("%f", &p)
	z := massacorp(a, p)
	fmt.Println(z)
	if z >= 25 {
		fmt.Println("precisa emagrecer")
	} else {
		fmt.Println("tudo OK")
	}
	fmt.Println(&a, a)
	fmt.Println(&p, p)
}

func massacorp(alx, pey float64) float64 {
	//z := alx * alx
	return pey / alx / alx
}
