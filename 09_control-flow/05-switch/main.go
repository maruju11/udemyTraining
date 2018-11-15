package main

import "fmt"

func main() {
	switch "maruju" {
	case "jorge":
		fmt.Println("jorge")
	case "maruju2":
		fmt.Println("Maruju2")
		fallthrough
	case "maruju":
		fmt.Println("Maruju")
	default:
		fmt.Println("sem amigos")
	}
}
