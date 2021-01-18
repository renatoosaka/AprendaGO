package main

import (
	"fmt"
)

func main() {
	pessoas := [][]string{
		[]string{
			"Joao", "Maradona", "Jogar Bola",
		},
		[]string{
			"Marilia", "Mendonça", "Cantar",
		},
		[]string{
			"Carlos", "Cunha", "Pilotar",
		},
	}

	for _, v := range pessoas {
		fmt.Println(v)
	}
}
