package main

import "fmt"

func main() {
	slice := [][]int{{10}, {20, 30}}
	slice[0] = append(slice[0], 11)
	fmt.Printf("%T  - %d  - %d\n", slice, len(slice), slice[0])
}
