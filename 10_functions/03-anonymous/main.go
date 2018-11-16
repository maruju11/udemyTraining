package main

import "fmt"

func greet() func() string {
	return func() string {
		return "Parabéns"
	}
}
func main() {
	hh := greet()
	fmt.Println(hh())
	fmt.Printf("%T \n", hh)
}
