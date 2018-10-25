package main

import "fmt"

func main() {
	v := 11
	b := incr(&v)
	fmt.Println(b)
	fmt.Println(incr(&v))
}

func incr(p *int) int {
	*p++
	return *p
}
