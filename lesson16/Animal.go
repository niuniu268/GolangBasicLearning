package main

import "fmt"

// essence is a point
type Animal interface {
	Sleep()
	GetColor() string
	GetType() string
}

// Specific class
type Cat struct {
	color string
}

func (this *Cat) Sleep() {

	fmt.Println("cat sleep")

}

func (this *Cat) GetColor() string {

	return this.color

}

func (this *Cat) GetType() string {

	return "CAT"

}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {

	fmt.Println("cat sleep")

}

func (this *Dog) GetColor() string {

	return this.color

}

func (this *Dog) GetType() string {

	return "Dog"

}
