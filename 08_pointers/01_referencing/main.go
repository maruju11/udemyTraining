package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a)
	fmt.Println(&a)

	var b = &a
	fmt.Println(b)
	fmt.Println("*b = ", *b)
	fmt.Println(&b)
	*b = 111
	fmt.Println("&a = ", &a)
	fmt.Println(a)
	fmt.Println(b)
	var x, y int
	fmt.Println(&x == &y, &x == &x, &x == nil)
}
