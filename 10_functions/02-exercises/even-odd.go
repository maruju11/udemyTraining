package main

import "fmt"

func pair(x int) bool {
	return x%2 == 0
}

func main() {
	fmt.Println(pair(3))
}
