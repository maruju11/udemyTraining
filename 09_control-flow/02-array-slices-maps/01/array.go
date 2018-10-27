package main

import "fmt"

func main() {
	x := [5]float64{1, 2, 3, 4, 500}
	tot := 0.0
	//x[3] = 100
	fmt.Println(x)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		tot = tot + x[i]
	}
	fmt.Println("total = ", tot)
}
