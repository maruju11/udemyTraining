package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0, 10, 5}

	c := make(chan int)
	go sum(s[:len(s)/4], c)
	go sum(s[len(s)/2:], c)
	d := len(s) / 3
	fmt.Println(d)
	//go sum(s[len(s)], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
