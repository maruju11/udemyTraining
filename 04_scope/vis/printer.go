package main

import "fmt"

var x = 42.0

func main() {
	fmt.Println(x)
	foo()
}
func foo() {
	fmt.Println(x * 2.1)
}
