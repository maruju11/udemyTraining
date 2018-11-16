package main

import "fmt"

func main() {
	x := filter([]int{1, 2, 3, 4, 5, 6, 73}, func(n int) bool {
		return n > 2
	})
	fmt.Println(x)
}

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	fmt.Println(xs)
	fmt.Printf("%T \n", xs)
	fmt.Printf("%T \n", callback)
	return xs
}
