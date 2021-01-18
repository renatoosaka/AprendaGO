package main

import (
	"fmt"
)

func main() {
	pessoas := map[string][]string{
		"brogodo_jose": []string{
			"pescar", "dormir",
		},
		"maradona_mario": []string{
			"jogar bola",
		},
		"silva_magali": []string{
			"comer", "brigar", "correr",
		},
	}

	for i, v := range pessoas {
		fmt.Println(i, v)
	}

	fmt.Println("----------------------")
	pessoas["fenomeno_ronaldo"] = []string{
		"jogar bola", "jogar poker",
	}

	for i, v := range pessoas {
		fmt.Println(i, v)
	}

	fmt.Println("----------------------")
	delete(pessoas, "silva_magali")

	for i, v := range pessoas {
		fmt.Println(i, v)
	}
}
