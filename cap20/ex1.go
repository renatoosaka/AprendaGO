package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go func() {
		for i := 0; i <= 15; i++ {
			fmt.Println("(func1)", i)
			time.Sleep(1000)
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i <= 15; i++ {
			fmt.Println("(func2)", i)
			time.Sleep(2000)
		}

		wg.Done()
	}()

	wg.Wait()
}
