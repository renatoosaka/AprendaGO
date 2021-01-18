package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 5, 10)

	slice[0], slice[1], slice[2], slice[3] = 1, 2, 3, 4

	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)

	fmt.Printf("%v %v %v\n", slice, len(slice), cap(slice))

	slice = append(slice, 11)

	fmt.Printf("%v %v %v\n", slice, len(slice), cap(slice))
}
