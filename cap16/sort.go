package main

import (
	"fmt"
	"sort"
)

func main() {
	ss := []string{"atirei", "o", "pau", "no", "gato"}

	si := []int{43, 123, 12, 7, 23, 22, 7, 9}

	fmt.Println(ss, si)

	sort.Strings(ss)
	sort.Ints(si)

	fmt.Println(ss, si)
}
