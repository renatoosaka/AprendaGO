package main

import (
	"fmt"
)

func main() {
	x := 10

	exp := func(y int) {
		fmt.Println(y * 10)
	}

	exp(x)
}
