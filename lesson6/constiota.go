package main

import (
	"fmt"
	"strconv"
)

const (
	BEIJING = iota
	SHANGHAI
	SHENZEH
)

const (
	a, b = iota*2 + 1, iota*2 + 2
	c, d
	e, f
)

func main() {

	fmt.Printf(strconv.Itoa(BEIJING))
	fmt.Printf(strconv.Itoa(SHANGHAI))
	fmt.Printf(strconv.Itoa(SHENZEH))

}
