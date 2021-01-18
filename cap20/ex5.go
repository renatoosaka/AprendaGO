package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var contador int64

func main() {
	var wg sync.WaitGroup

	contador = 0
	wgCount := 10

	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("GoRoutine (1)", runtime.NumGoroutine())

	wg.Add(wgCount)

	for i := 0; i < wgCount; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()

			fmt.Println("Contador:", atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(contador)
	fmt.Println("GoRoutine (3)", runtime.NumGoroutine())
}
