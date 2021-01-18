package main

import (
	"fmt"
)

func main() {
	anonimo := struct {
		nome  string
		idade int
	}{"mario", 32}

	fmt.Println(anonimo)
}
