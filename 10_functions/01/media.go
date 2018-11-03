package 01

import "fmt"

func Media(xs []float64) float64 {
	total := 0.0
	for _, soma := range xs {
		total += soma
	}
	fmt.Println(total / float64(len(xs)))
}
