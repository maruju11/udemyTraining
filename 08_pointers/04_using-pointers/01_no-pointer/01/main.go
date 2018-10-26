package main

import "fmt"

func zz(z *int) {
	fmt.Println(z)
	fmt.Printf("%p\n", &z)
	*z = 0
	fmt.Println(&z)
}

func main() {
	x := 55
	fmt.Printf("%p\n", &x)
	zz(&x)
	fmt.Println(x)
	fmt.Println(&x)
}
