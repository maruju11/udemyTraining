package main

import "fmt"

func main() {
	fmt.Println([]byte("Maruju"))

	for i := 50; i <= 140; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	foo := 'a'
	too := "a"
	fmt.Println(foo, " - ", too)
	fmt.Printf("%T ", foo)
	fmt.Printf("%T \n", too)
	fmt.Printf("%T ", rune(foo))
	fmt.Printf("%T \n", rune(foo))
}
