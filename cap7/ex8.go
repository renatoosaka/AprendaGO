package main

import (
	"fmt"
)

func main() {
	year := 2020

	switch {
	case year == 2020:
		{
			fmt.Println("Ano da Covid-19")
		}
	case year == 2021:
		{
			fmt.Println("Continuação da Covid-19")
		}
	}
}
