package main

import (
	"fmt"
)

func main() {
	funcao := retornaFuncao()

	fmt.Println(funcao())
}

func retornaFuncao() func() string {
	return func() string {
		return "ola mundo"
	}
}
