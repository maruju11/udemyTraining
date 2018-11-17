package main

import "fmt"

func main() {
	//var m [4][2]int
	var n = [4][2]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}}
	fmt.Println(n)
	//m = n
	mm(&n)
	//fmt.Println(n)
	//fmt.Println(&n)
	fmt.Println(n[1][1])
}

func mm(x *[4][2]int) {
	fmt.Println(x)
	fmt.Println(*x)
}
