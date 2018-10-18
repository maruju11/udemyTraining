package main

import (
	"fmt"
)

func main() {
	const boilF, freezF = 212.0, 32.0

	fmt.Printf("temperatura de fervura = %goF ou %goC\n", f, c)
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
