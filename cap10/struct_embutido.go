package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {
	pessoa1 := pessoa{"Alfredo", 30}
	pessoa2 := profissional{pessoa{"Maricota", 31}, "Pizzaiola", 10000}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
}
