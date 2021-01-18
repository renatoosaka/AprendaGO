package main

import (
	"fmt"
)

func main() {
	year := 1983

	for {
		if year > 2021 {
			break
		}

		fmt.Println(year)
		year++
	}
}
