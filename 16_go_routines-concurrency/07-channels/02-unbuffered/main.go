package main

import (
	"fmt"
)

func main() {
	//unbuffered channel
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	//time.Sleep(time.Second)

}
