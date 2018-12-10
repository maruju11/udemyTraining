package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"solange", "Adriana", "Mila", "Rita", "Sandra"}
	fmt.Println(s)
	sort.StringSlice(s).Sort()
	//sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
}
