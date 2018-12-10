package main

import "fmt"
import "sort"

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 22, 32, 3, 11, 89, 102}
	fmt.Println(n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}
