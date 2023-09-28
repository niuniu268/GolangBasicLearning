package main

import (
	"fmt"
	//"runtime"

	"time"
)

func main() {

	go func() {

		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")

			fmt.Println("B")

			//runtime.Goexit()

		}()
		fmt.Println("A")

	}()

	for true {
		time.Sleep(10 * time.Second)
	}

}
