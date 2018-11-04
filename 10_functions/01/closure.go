package main

import (
	"fmt"
)

func main() {
	//tot := 0
	add := func(x, y, z, w int) int {
		return x + y + z + w
	}
	fmt.Println(add(1, 2, 3, 4))
	fmt.Println(fatorial(30))
}

func fatorial(xx uint) uint {
	if xx == 0 {
		return 1
	}
	return xx * fatorial(xx-1)
}
