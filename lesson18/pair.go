package main

import "fmt"

func main() {

	var a string
	a = "test"

	//	a pair<type:string, value:"test>

	// allType pair<type:string, value:"test">

	var allType interface{}
	allType = a

	str, _ := allType.(string)

	fmt.Println(str)

}
