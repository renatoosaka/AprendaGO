package main

import (
	"fmt"
)

func main() {
	callback(func() {
		fmt.Println("callback")
	})
}

func callback(c func()) {
	c()
}
