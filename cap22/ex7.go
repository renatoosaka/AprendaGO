package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	for i := 0; i < 10; i++ {
		go func(c chan int) {
			for j := 0; j < 10; j++ {
				c <- j * i
				time.Sleep(time.Duration(time.Second))
			}
		}(c)
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

}
