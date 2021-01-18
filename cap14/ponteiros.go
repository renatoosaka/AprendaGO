package main

import (
	"fmt"
)

func main() {
	x := 10

	y := &x

	fmt.Println(x, &x)
	fmt.Println(y, *y)
	fmt.Printf("%T %T\n", x, y)

	a := 0
	b := &a

	fmt.Println(a, *b)

	*b++

	fmt.Println(a, *b)

	z := 0
	fmt.Println(z)

	valor(z)
	ponteiro(&z)

	fmt.Println(z)
}

func valor(v int) {
	v++
	fmt.Println("(valor)", v)
}

func ponteiro(p *int) {
	*p++
	fmt.Println("(ponteiro)", *p)
}
