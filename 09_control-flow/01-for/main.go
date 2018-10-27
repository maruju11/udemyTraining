package main

import "fmt"

func main() {
	i := 1
	for i <= 21 {
		fmt.Println("valor de i = ", i)
		if i%3 == 0 {
			fmt.Println(i, "even")
		} //else {
		//fmt.Println(i, "odd")
		//}
		i++
	}
}
