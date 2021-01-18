package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) souEu() {
	fmt.Println("Meu nome é", p.nome, p.sobrenome, "e tenho", p.idade, "anos")
}

func main() {
	eu := pessoa{"Renato", "Osaka", 37}

	eu.souEu()
}
