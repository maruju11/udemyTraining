package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ { // diz respeito ao numero de caraacteres w palavras
		s += sep + os.Args[i]
		fmt.Println(len(os.Args[i]))
		fmt.Printf("%d \t - Args = %s\n", i, s)
		sep = "x"
	}
	fmt.Println(os.Args)
	fmt.Println(s)
}
