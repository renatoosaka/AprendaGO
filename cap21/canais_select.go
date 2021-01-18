package main

import (
	"fmt"
)

func main() {
	canal := make(chan int)
	quit := make(chan int)
	a := make(chan int)
	b := make(chan int)

	x := 500

	go func(i int) {
		for x := 0; x < i; x++ {
			a <- x
		}
	}(x / 2)

	go func(i int) {
		for x := 0; x < i; x++ {
			b <- x
		}
	}(x / 2)

	for i := 0; i < x; i++ {
		select {
		case v := <-a:
			fmt.Println("Canal A:", v)
		case v := <-b:
			fmt.Println("Canal B:", v)
		}
	}
	fmt.Println("")

	go recebeQuit(canal, quit)
	enviaPraCanal(canal, quit)
}

func recebeQuit(canal chan int, quit chan int) {
	for i := 0; i < 500; i++ {
		fmt.Println("recebido", <-canal)
	}
	quit <- 0
}

func enviaPraCanal(canal chan int, quit chan int) {
	qualquercoisa := 10

	for {
		select {
		case canal <- qualquercoisa:
			qualquercoisa++
		case <-quit:
			return
		}
	}
}
