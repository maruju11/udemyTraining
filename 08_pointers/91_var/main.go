package main

import "fmt"

func quad(x *float64) {
	*x = *x * *x
}

func main() {
	x := 30.0
	quad(x)
	fmt.Println(&x)
	fmt.Println(x)
	fmt.Println(222)
}
