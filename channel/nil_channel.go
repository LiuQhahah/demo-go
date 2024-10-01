package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for v := range asChan(2, 3, 3, 4, 6, 4, 4, 4, 4) {
		fmt.Printf("output: %d\n", v)
	}

}

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			fmt.Printf("v:%d\n", v)
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()

	return c
}
