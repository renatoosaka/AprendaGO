package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)

	go send(canal)
	receive(canal)
}

func send(s chan<- int) {
	s <- 32
}

func receive(r <-chan int) {
	fmt.Println("Valor do canal", <-r)
}
