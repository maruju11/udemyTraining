package main

import "fmt"

const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
	d = iota // 0
	e = iota // 0
	f = iota // 0
)
const (
	_  = iota             // 0
	kB = 1 << (iota * 10) // 1<<(1*10))
	mB = 1 << (iota * 10) // 1<< (2*10)
	gB = 1 << (iota * 10) // 1 << (3 * 10)
)

func main() {
	fmt.Println(a, b, c, d, e, f)
	fmt.Printf("binário  \t\tdecimal\n")
	fmt.Printf("%b\t", kB)
	fmt.Printf("%d \t\n", kB)
	fmt.Printf("%b\t", mB)
	fmt.Printf("%d\t\n", mB)
	fmt.Printf("%b\t", gB)
	fmt.Printf("%d\t\n", gB)
}
