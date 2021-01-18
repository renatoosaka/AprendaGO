package main

import (
	"fmt"
)

func main() {
	fmt.Println(retornaInt())

	n, s := retornaIntStr()

	fmt.Println(n, s)
	fmt.Println(retornaIntStr())
}

func retornaInt() int {
	return 100
}

func retornaIntStr() (int, string) {
	return 10, "dez"
}
