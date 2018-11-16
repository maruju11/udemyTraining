package main

import (
	"fmt"
)

func main() {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := maior(data...)
	m := media(data...)
	fmt.Println(n, m)
}
func maior(args ...float64) float64 {
	tot := 0.0
	for _, v := range args {
		if tot <= v {
			tot = v
		}
	}
	return tot
}
func media(args1 ...float64) float64 {
	tot1 := 0.0
	for _, v1 := range args1 {
		tot1 += v1
	}
	fmt.Println(args1)
	fmt.Println(len(args1))
	fmt.Println(tot1)
	return tot1 / float64(len(args1))

}
