package main

import "fmt"

func Demo(i int) {
	var arr [10]int

	defer func() {
		// set up recover in order intercept error message
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()
	// imitate a mistake
	arr[i] = 10
}

func main() {

	Demo(10)

	fmt.Println("continue process")

}
