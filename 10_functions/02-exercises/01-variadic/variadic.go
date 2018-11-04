package main

import (
	"fmt"
)

func maior(args ...int) int {
	tot := 0
	for _, v := range args {
		if tot <= v {
			tot = v
		}

	}
	return tot
}

func main() {
	fmt.Println(maior(1, 2, 22, 3, 4, 5, 6, 7))
}
