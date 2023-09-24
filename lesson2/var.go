package main

import "fmt"

//the fourth type only is used under a method

func main() {
	// the first type of variable declaration
	var a int = 10
	fmt.Println("a= ", a)
	fmt.Printf("type of a is %T\n", a)

	// the second type of variable declaration
	var b int
	b = 100
	fmt.Println("b= ", b)
	fmt.Printf("type of b is %T\n", b)

	// the third type of variable declaration

	var c = 1000

	fmt.Println("c= ", c)
	fmt.Printf("type of c is %T\n", c)

	// the fourth type of variable declaration

	e := 1000

	fmt.Println("e= ", e)
	fmt.Printf("type of e is %T\n", e)

	// declare multiple variables

	var kk, ll = 100, "variable"
	fmt.Println("kk = ", kk, "ll = ", ll)

	var (
		aa = 11
		bb = 22
	)
	fmt.Println("aa = ", aa, "bb = ", bb)

}
