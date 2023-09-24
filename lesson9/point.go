package main

import "fmt"

func Swap(a string, b string) {

	tmp := ""

	tmp = a
	a = b
	b = tmp

}
func SwapPlace(a *string, b *string) {

	var tmp = *a
	*a = *b
	*b = tmp

}

func main() {

	a := "a"
	b := "b"

	fmt.Println("a is ", a)
	fmt.Println("b is ", b)

	Swap(a, b)

	fmt.Println("a is ", a)
	fmt.Println("b is ", b)

	SwapPlace(&a, &b)

	fmt.Println("a is ", a)
	fmt.Println("b is ", b)

	var p *string

	p = &a
	fmt.Println("p is ", p)
	fmt.Println("&a is ", &a)

	var pp **string
	pp = &p

	fmt.Println("&p is ", &p)
	fmt.Println("pp is ", pp)

}
