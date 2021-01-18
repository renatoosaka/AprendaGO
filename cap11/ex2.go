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
	pessoas := map[string]pessoa{
		"osaka":   pessoa{"Renato", "Osaka", []string{"chocolate", "morango"}},
		"pereira": pessoa{"Vanessa", "Osaka", []string{"doce de leite", "chocolate", "pistache"}},
	}

	for _, p := range pessoas {
		fmt.Println(p)

		for _, v := range p.sorvetes_favoritos {
			fmt.Println("\t", v)
		}
	}

}
