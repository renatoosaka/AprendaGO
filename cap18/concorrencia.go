package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("Go Routines (1):", runtime.NumGoroutine())

	go func1()
	go func2()

	fmt.Println("Go Routines (2):", runtime.NumGoroutine())

	wg.Wait()
}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1:", i)
		time.Sleep(20)
	}

	wg.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func2:", i)
	}

	wg.Done()
}
