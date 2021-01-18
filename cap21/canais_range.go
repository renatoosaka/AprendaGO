package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)

	go meuLoop(10, canal)
	prints(canal)
}

func meuLoop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i
	}

	close(s)
}

func prints(r <-chan int) {
	for v := range r {
		fmt.Println("From channel", v)
	}
}
