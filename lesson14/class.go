package main

import "fmt"

func main() {

	a := Animal{age: 1, name: "cat", color: "white"}

	fmt.Println("before method1", a.name)

	a.Scream1()

	fmt.Println("after method1", a.name)

	fmt.Println("before method2", a.name)

	a.Scream2()

	fmt.Println("after method2", a.name)

	//a.Setname1("wangwang")
	fmt.Println("getname", a.Getname())

	a.Setname1("wangwang")
	fmt.Println("setname without point", a.Getname())
	a.Setname2("awuawu")

	fmt.Println("setname with point", a.Getname())

}
