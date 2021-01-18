package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

func main() {
	eu := pessoa{"Renato", 37}
	fmt.Println(eu)

	mudeMe(&eu)
	fmt.Println(eu)
}

func mudeMe(p *pessoa) {
	(*p).nome = fmt.Sprintf("%v mudou", (*p).nome)
	(*p).idade = (*p).idade * 2
}
