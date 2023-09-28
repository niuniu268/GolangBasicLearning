package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		defer fmt.Println("goroute is end")
		fmt.Println("goroute is runing")

		// send 666 to c
		c <- 666
	}()

	// num receive 666 from c
	num := <-c

	fmt.Println("num's value is ", num)

}
