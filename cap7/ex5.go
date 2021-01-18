package main

import (
	"fmt"
)

func main() {
	for x := 10; x <= 100; x++ {
		fmt.Printf("%v %v\n", x, (x % 4))
	}
}
