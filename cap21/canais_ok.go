package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)

	go func() {
		canal <- 0
		close(canal)
	}()

	v, ok := <-canal

	fmt.Println(v, ok)

	v, ok = <-canal

	fmt.Println(v, ok)

	par := make(chan int)
	ímpar := make(chan int)
	quit := make(chan bool)

	go mandaNúmeros(par, ímpar, quit)

	receive(par, ímpar, quit)
}

func mandaNúmeros(par, ímpar chan int, quit chan bool) {
	total := 100
	for i := 0; i < total; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			ímpar <- i
		}
	}
	close(par)
	close(ímpar)
	quit <- true
}

func receive(par, ímpar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("O número", v, "é par.")
		case v := <-ímpar:
			fmt.Println("O número", v, "é ímpar.")
		case v, ok := <-quit:
			if !ok {
				fmt.Println("Deu zebra! Ó:", v)
			} else {
				fmt.Println("Encerrando. Recebemos:", v)
			}
			return
		}
	}
}
