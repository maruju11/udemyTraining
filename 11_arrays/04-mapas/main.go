package main

import "fmt"

func main() {
	colors := map[string]string{}
	colors["red"] = "#da1337"
	colors["blue"] = "#ff1234"
	colors["coral"] = "#a90223"
	fmt.Println(colors)
	for key, value := range colors {
		fmt.Printf("key = %s - value=%s\n", key, value)
	}
	value, exists := colors["blue"]

	if exists {
		fmt.Printf("%s\n", value)
	} else {
		fmt.Println("nao existe")
	}
}
