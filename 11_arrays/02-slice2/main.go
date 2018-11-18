package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	newSlice := slice[2:3:4]
	newSlice = append(newSlice, 6)
	newSlice = append(newSlice, 60)
	newSlice = append(newSlice, 160, 131, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	slice = append(slice, 222, 231, 21, 22, 3, 4, 5, 6, 7, 8, 22, 23, 24, 25, 25)
	fmt.Println(slice, newSlice)
	//fmt.Printf("%T \n", slice)
	fmt.Printf("%T \n", newSlice)
	fmt.Println(len(slice), len(newSlice))
}
