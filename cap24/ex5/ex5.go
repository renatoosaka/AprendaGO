package main

import (
	"fmt"
)

func main() {
	avg := Average([]float64{1, 2, 3, 4, 5, 6})

	fmt.Println(avg)
}

func Average(numbers []float64) float64 {
	var t float64
	t = 0

	for _, v := range numbers {
		t += v
	}

	return t / float64(len(numbers))
}
