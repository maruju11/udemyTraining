package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	newSlice := slice[0:2:5]
	fmt.Println(newSlice)
	newSlice = append(newSlice, 6)
	newSlice = append(newSlice, 7)
	fmt.Println(newSlice)
	newSlice = append(newSlice, 8, 9, 10, 11, 12, 13)
	slice = append(slice, 222, 231, 21, 22, 23, 24)
	fmt.Println(slice, newSlice)
	//fmt.Printf("%T \n", slice)
	fmt.Printf("%T \n", newSlice)
	fmt.Println(len(slice), len(newSlice))
}
