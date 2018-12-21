package main

import "fmt"

func main() {
	m := 4
	f := factorial(m)
	for n := range f {
		fmt.Println(n)
	}
}

/* func factorial(a int) int {
	total := 1
	for i := a; i > -0; i-- {
		total *= i
	}
	return total
}
*/
func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
