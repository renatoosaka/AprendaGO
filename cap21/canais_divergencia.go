package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go manda(20, c1)
	go outra(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
}

func manda(n int, c chan int) {
	defer close(c)

	for i := 0; i < n; i++ {
		c <- i
	}
}

func outra(c1, c2 chan int) {
	var wg sync.WaitGroup

	defer close(c2)

	for v := range c1 {
		wg.Add(1)

		go func(x int) {
			c2 <- trabalho(x)
			wg.Done()
		}(v)
	}

	wg.Wait()
}

func trabalho(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return n * rand.Intn(1e3)
}
