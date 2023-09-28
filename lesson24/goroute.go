package main

import (
	"fmt"
	"time"
)

func NewTask() {
	i := 0

	for true {
		i++
		fmt.Printf("newtask in goroute: i %d\n", i)

		time.Sleep(1 * time.Second)

	}

}

func main() {

	go NewTask()

	i := 0

	for true {
		i++

		fmt.Printf("main in goroute: i %d\n", i)

		time.Sleep(1 * time.Second)
	}
}
