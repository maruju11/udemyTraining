package main

import "fmt"

const (
	a = iota // 0
	B = iota // 1
	C = iota // 2
	D = iota // 0
	E = iota // 0
	F = iota // 0
)
const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1<<(1*10))
	MB = 1 << (iota * 10) // 1<< (2*10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
)

func main() {
	fmt.Println(a, B, C, D, E, F)
	fmt.Printf("binário  \t\tdecimal\n")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d \t\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\t\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\t\n", GB)
}
