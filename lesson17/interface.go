package main

import "fmt"

func myFun(arg interface{}) {

	fmt.Println("myfun is called")
	fmt.Println(arg)

	//	declaration interface{}

	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string")
	} else {
		fmt.Println("arg is string, value", value)
		fmt.Printf("value type is %T\n", value)
	}

}

type Book struct {
	auth string
}

func main() {
	var book Book
	book.auth = "golang"
	myFun(book)
	myFun(100)
	myFun("test")

}
