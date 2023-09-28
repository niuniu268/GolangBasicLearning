package main

import (
	"fmt"
	"reflect"
)

type User struct {
	id   string
	name string
	age  int
}

func (this *User) Call() {
	fmt.Println("func call")

	fmt.Printf("%v\n", this)

}

func main() {

	a := &User{
		id:   "2",
		name: "miaomiao",
		age:  2,
	}

	a.Call()

	GetFieldsandMethods(a)

}

func GetFieldsandMethods(arg interface{}) {

	numMethod := reflect.TypeOf(arg).NumMethod()
	fmt.Println(numMethod)

	for i := 0; i < numMethod; i++ {

		fmt.Println(reflect.TypeOf(arg).Method(i).Name)
		fmt.Println(reflect.TypeOf(arg).Method(i).Type)
		fmt.Println(reflect.TypeOf(arg).Method(i).Func)

	}

	reflect.ValueOf(arg)
	fmt.Println(reflect.ValueOf(arg))

	numField := reflect.ValueOf(arg).Elem().NumField()

	fmt.Println(numField)
	for i := 0; i < numField; i++ {

		fmt.Println(reflect.ValueOf(arg).Elem().Field(i))

	}

}
