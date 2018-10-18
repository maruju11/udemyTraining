package main

import "fmt"

func main() {
	a := 100.0
	fmt.Println("a = ", a)
	fmt.Println("endereço de a - ", &a)
	fmt.Printf("%d \n", &a)
	indM()
}

func indM() float64 {
	var pesoKg float64
	var altura float64
	fmt.Print("digite seu peso: = ")
	fmt.Scan(&pesoKg)
	fmt.Print("digite sua altura =")
	fmt.Scan(&altura)
	resultado := pesoKg / (altura * altura)
	fmt.Println(resultado)

}
