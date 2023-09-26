package main

import "fmt"

type Reader interface {
	ReadBook()
}
type Writer interface {
	WriteBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Readbook is used")

}

func (this *Book) WriteBook() {
	fmt.Printf("Writebooks is used")

}

func main() {
	//<type point, value &Book{}>
	b := &Book{}
	// <type nil, value nil>
	var r Reader
	fmt.Printf("&T\n", r)

	//<type point, value &Book>
	r = b
	fmt.Printf("&T\n", r)
	r.ReadBook()

	var w Writer
	fmt.Printf("&T\n", w)

	w = b
	fmt.Printf("&T\n", w)
	w = r.(Writer)
	fmt.Printf("&T\n", w)
	w.WriteBook()

}
