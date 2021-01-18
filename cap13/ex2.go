package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum1([]int{1, 2, 3}...))
	fmt.Println(sum2([]int{1, 2, 3}))
}

func sum1(n ...int) int {
	total := 0

	for _, v := range n {
		total += v
	}

	return total
}

func sum2(n []int) int {
	total := 0

	for _, v := range n {
		total += v
	}

	return total
}
