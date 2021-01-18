package main

import (
	"fmt"
)

func main() {
	anonimo := struct {
		mapa   map[string]string
		islice []string
	}{
		mapa: map[string]string{
			"SP": "São Paulo",
			"RJ": "Rio de Janeiro",
		},
		islice: []string{"Janeiro", "Fevereiro", "Março"},
	}

	fmt.Println(anonimo)
	fmt.Println(anonimo.mapa["SP"])
}
