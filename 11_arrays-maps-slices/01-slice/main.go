package main

import "fmt"

func main() {
	slice1 := make([]int, 3, 5)
	slice2 := []int{1, 2, 3}
	slice3 := []int{9: 0}
	array1 := [3]int{4, 5, 6}
	fmt.Printf("%T = slice \n", slice1)
	fmt.Printf("%T = slice2 \n", slice2)
	fmt.Printf("%X = slice3  len=%d   capac= %d\n", &slice3, len(slice3), cap(slice3))
	fmt.Printf("%T = array \n", array1)
}
