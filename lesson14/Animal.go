package main

import "fmt"

type Animal struct {
	color string
	name  string
	age   int
}

func (this Animal) Scream1() {

	this.name = "miaomiao"

	fmt.Println(this.name)

}

func (this *Animal) Scream2() {
	this.name = "miaomiao"

	fmt.Println(this.name)

}

func (this Animal) Getname() string {

	return this.name

}

func (this Animal) Setname1(newname string) {

	this.name = newname

}

func (this *Animal) Setname2(newname string) {

	this.name = newname

}
