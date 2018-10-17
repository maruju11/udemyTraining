package main

import (
	"fmt"
)

const boilF = 212.0

func main() {
	var f = boilF
	var c = (f - 32) * 5 / 9
	fmt.Printf("temperatura de fervura = %goF ou %goC\n", f, c)
}
