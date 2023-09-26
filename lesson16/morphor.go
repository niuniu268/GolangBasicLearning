package main

import "fmt"

func main() {

	var a Animal

	a = &Cat{
		color: "white",
	}

	fmt.Println(a.GetType())
	fmt.Println(a.GetColor())

	a.Sleep()

	a = &Dog{
		color: "black",
	}

	fmt.Println(a.GetType())
	fmt.Println(a.GetColor())

	a.Sleep()

}
