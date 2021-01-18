package main

import (
	"fmt"
)

type pessoa struct {
	nome               string
	sobrenome          string
	sorvetes_favoritos []string
}

func main() {
	pessoa1 := pessoa{"Renato", "Osaka", []string{"chocolate", "morango"}}
	pessoa2 := pessoa{"Vanessa", "Osaka", []string{"doce de leite", "chocolate", "pistache"}}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)

	for _, v := range pessoa1.sorvetes_favoritos {
		fmt.Println(v)
	}
}
