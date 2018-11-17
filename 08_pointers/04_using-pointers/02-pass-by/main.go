package main

import "fmt"

func main() {
	maruju := 11
	fmt.Println(&maruju)
	fmt.Println(maruju)
	altMar(&maruju)
	fmt.Println("main2 = ", &maruju)
	fmt.Println(maruju)
}

func altMar(x *int) {
	fmt.Println("x = ", x)
	fmt.Println(*x)
	*x = 1111111
	fmt.Println("x = 111 new", x)
	fmt.Println(*x)
	fmt.Println(&x)
}
