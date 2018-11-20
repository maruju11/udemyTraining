package main

import "fmt"

func main() {
	slice1 := make([]int, 2, 5)
	fmt.Println(slice1)
	for i := 0; i <= 80; i++ {
		slice1 = append(slice1, i)
		fmt.Println("lin = ", len(slice1), "cap=", cap(slice1), "valor=", slice1[i])
	}
}
