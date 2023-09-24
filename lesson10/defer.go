package main

import "fmt"

func Test() {
	defer fmt.Println("defer comment in method")

	fmt.Println("normal comment imitate return")

}
func main() {
	defer fmt.Println("defer comment")
	fmt.Println("normal comment")

	Test()

}
