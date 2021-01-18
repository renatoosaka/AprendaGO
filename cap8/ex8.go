package main

import (
	"fmt"
)

func main() {
	pessoas := map[string]string{
		"brogodo_jose":   "pescar",
		"maradona_mario": "jogar bola",
		"silva_magali":   "comer",
	}

	for i, v := range pessoas {
		fmt.Println(i, v)
	}
}
