package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 3)
	fmt.Println("channel length", len(c), "capability ", cap(c))

	go func() {
		fmt.Println("goroute start")

		for i := 0; i < 4; i++ {

			c <- i
			fmt.Println("goroute send element i= ", i)

		}
	}()

	defer fmt.Println("main is end")

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {

		num := <-c

		fmt.Println("num receive the channel i, num= ", num)

	}
}
