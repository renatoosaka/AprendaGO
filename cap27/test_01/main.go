package example

import (
	"fmt"
)

func main() {
	fmt.Println(Sum(1, 2, 3, 4))
}

//Sum soma diversos valores
func Sum(i ...int) int {
	t := 0

	for _, v := range i {
		t += v
	}

	return t
}
