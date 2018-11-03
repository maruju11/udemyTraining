package main

import (
	"errors"
	"fmt"
	"os"
)

func Media(xs []float64) float64 {
	total := 0.0
	for _, soma := range xs {
		fmt.Print(soma)
		total += soma
	}
	fmt.Println(xs)
	ff(111, 112)
	return total / float64(len(xs)) //fmt.Println(total / float64(len(xs)))

}

func ff(p1, p2 int) {
	fmt.Println(p1, p2)
}

func main() {
	x := []float64{110, 120, 130, 140, 150, 160, 170, 180}
	fmt.Println(Media(x))
	var CommandLine = NewFlagSet(os.Args[0], ExitOnError)
	var ErrHelp = errors.New("flag: help requested")
	var Usage = func() {
		fmt.Printf(CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		PrintDefaults()
	}
}
