package main

import "fmt"

func main() {
	var z []float64
	fmt.Println(z)
	y := make([]float64, 010)
	fmt.Println("y = ", y[5:6])
	slice1 := []float64{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Print(slice1, slice2, "\n")

	copy(slice2, y)
	fmt.Println("copy =", slice2, y)

	arr := [5]float64{1, 2, 3, 4, 5}
	x := arr[1:len(arr)]
	fmt.Println(x)
}
