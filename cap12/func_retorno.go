package main

import (
	"fmt"
)

func main() {
	funcao1 := retornaFuncao1(10)

	fmt.Println(funcao1())
	fmt.Println(retornaFuncao1(4)())

	funcao2 := retornaFuncao2(10)

	fmt.Println(funcao2(2))

	fmt.Println(retornaFuncao2(10)(5))
}

func retornaFuncao1(x int) func() int {
	return func() int {
		return x * 10
	}
}

func retornaFuncao2(x int) func(int) int {
	return func(b int) int {
		return x * 10 * b
	}
}
