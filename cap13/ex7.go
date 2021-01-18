package main

import (
	"fmt"
)

func main() {
	funcao := func(s string) {
		fmt.Println(s)
	}

	funcao("hello world")
}
