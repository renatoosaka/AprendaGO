package main

import (
	"fmt"
)

type batata int

var x batata

func main() {
	fmt.Printf("%v %T\n", x, x)

	x = 42

	fmt.Printf("%v %T\n", x, x)

}
