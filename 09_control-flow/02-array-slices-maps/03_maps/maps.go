package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 11223344
	x["jorge"] = 111
	x["maruju"] = 1111
	if name, ok := x["maruj"]; ok {
		fmt.Println(name, ok)
	}
	fmt.Println(x["key"])
}
