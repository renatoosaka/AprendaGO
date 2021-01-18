package main

import (
	"fmt"
)

func main() {
	valor := soma(1, 2)
	fmt.Println(valor)

	hello()
	hi("Renato")

	total1, qtd1 := total(1, 2, 3)
	fmt.Println(total1, qtd1)

	numbers := []int{5, 6, 7}
	total2, qtd2 := total(numbers...)
	fmt.Println(total2, qtd2)
}

func hello() {
	fmt.Println("Hello")
}

func hi(name string) {
	fmt.Println("Hi", name)
}

func soma(x, y int) int {
	return x + y
}

func total(x ...int) (int, int) {
	sum := 0

	for _, v := range x {
		sum += v
	}

	return sum, len(x)
}
