package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i

			fmt.Println("goroute is runing, i= ", i)

			// channel cannot restart

			//close(c)
		}

		//close channel c

		close(c)
	}()

	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("main is finished")

}
