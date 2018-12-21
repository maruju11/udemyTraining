package main

import "fmt"

func main() {
	c := make(chan int)
	// c <- 1 - se deixar esta linha = deadlock para resolver, precisamos de go func
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}
