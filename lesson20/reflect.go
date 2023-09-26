package main

import (
	"fmt"
	"reflect"
)

func ReflectNum(arg interface{}) {
	reflect.TypeOf(arg)
	reflect.ValueOf(arg)

	fmt.Println("num's type is", reflect.TypeOf(arg))
	fmt.Println("num's value is ", reflect.ValueOf(arg))

}

func main() {

	var num float64 = 1.2345

	ReflectNum(num)

}
