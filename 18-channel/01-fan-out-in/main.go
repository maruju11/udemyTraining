package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3, 4, 5, 6, 7, 8, 9, 10)
	in2 := gen(11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	c1 := sq(in)
	c2 := sq(in)
	c3 := sq(in2)
	c4 := sq(in2)
	for n := range merge(c1, c2, c3, c4) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	fmt.Printf("Type NUms: %T\n", nums)
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return (out)
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	fmt.Printf("TYPE: %T\n", cs)
	wg.Add(len(cs))
	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
