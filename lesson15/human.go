package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("eating")
}
func (this *Human) work() {
	fmt.Println("working")
}

// inherit from human
type Superman struct {
	Human

	level int
}

// redefine parents method

func (this *Superman) Eat() {
	fmt.Println("absorbbing")
}

func (this *Superman) Fly() {
	fmt.Println("flying")
}
