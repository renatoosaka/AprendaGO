package main

import (
	"fmt"
)

func main() {
	a := closure()
	a()
	a()
	a()

	b := closure()
	b()
	b()
	b()
}

func closure() func() {
	x := 0

	return func() {
		x++
		fmt.Println(x)
	}
}
