package main

import (
	"fmt"
	"sync"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)

	go envia(par, impar)
	go recebe(par, impar, converge)

	for v := range converge {
		fmt.Println("(converge)", v)
	}
}

func envia(par, impar chan int) {
	defer close(par)
	defer close(impar)

	x := 100

	for i := 0; i < x; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
}

func recebe(par, impar, converge chan int) {
	var wg sync.WaitGroup

	defer close(converge)
	defer wg.Wait()

	wg.Add(2)

	go func() {
		defer wg.Done()

		for v := range par {
			converge <- v
		}
	}()

	go func() {
		defer wg.Done()

		for v := range impar {
			converge <- v
		}
	}()
}
