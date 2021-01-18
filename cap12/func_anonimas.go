package main

import (
	"fmt"
)

func main() {
	x := 10

	func(y int) {
		fmt.Println(y * 10)
	}(x)
}
