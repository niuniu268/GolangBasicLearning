package main

import (
	"fmt"
	"strconv"
)

func Test1() int {

	c := 100
	return c

}

func Test2() (int, int) {

	return 100, 200
}

func Test3() (r1, r2 int) {

	r1 = 1
	r2 = 3

	return

}

func main() {
	c := Test1()
	fmt.Println(strconv.Itoa(c))

	e, f := Test2()

	fmt.Println("e is", e)
	fmt.Println("f is", f)

	g, h := Test3()

	fmt.Println("g is ", g)
	fmt.Println("h is ", h)

}
