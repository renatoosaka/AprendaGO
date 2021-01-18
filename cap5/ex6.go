package main

import (
	"fmt"
)

const (
	a = iota + 2021
	b
	c
	d
)

func main() {
	fmt.Printf("%v\n%v\n%v\n%v", a, b, c, d)
}
