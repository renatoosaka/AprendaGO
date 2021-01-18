package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) bomDia() {
	fmt.Println(p.nome, "oi bom dia")
}

func main() {
	eu := pessoa{"Renato", 37}

	eu.bomDia()
}
