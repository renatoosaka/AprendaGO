package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("depois")
	fmt.Println("antes")
}
