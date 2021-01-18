package main

import (
	"fmt"
)

func main() {
	x := 11

	if x < 10 {
		fmt.Println("menor que deix")
	} else if x == 10 {
		fmt.Println("igual a deix")
	} else {
		fmt.Println("maior que deix")
	}
}
