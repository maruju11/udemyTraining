package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	newSlice := slice[0:3:3]
	fmt.Println(newSlice)
	newSlice = append(newSlice, 6)
	newSlice = append(newSlice, 7)
	fmt.Println(newSlice)
	newSlice = append(newSlice, 8, 9, 10)
	slice = append(slice, 21, 22, 23, 24)
	fmt.Println(slice, newSlice)
	//fmt.Printf("%T \n", slice)
	fmt.Printf("%T \n", newSlice)
	fmt.Println(len(slice), len(newSlice))
	fmt.Printf("%p\n", append(newSlice, slice...))
	for index, value := range slice {
		fmt.Printf("iNDEX = %d   VAlour: %d value-address= %X Elem-address = %X\n", index, value, &value, &slice[index])
	}
}
