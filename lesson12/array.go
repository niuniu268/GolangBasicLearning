package main

import "fmt"

var arr0 [4]int
var arr1 = []int{1, 2, 3}
var arr2 []int

var slice0 = make([]int, 3, 5)

func main() {

	//arr2[0] = 1

	arr2 = append(arr2, 0)
	arr2 = append(arr2, 2, 3)

	arr4 := make([]int, len(arr2), len(arr2)*2)

	fmt.Println(arr0)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(slice0)
	fmt.Println(arr4, " len", len(arr4), " cap", cap(arr4))
	var slice1 []int
	slice1 = make([]int, 6)
	copy(slice1, arr1)

	fmt.Println(slice1)

	for i := range slice1 {

		fmt.Println("print i ", i)

	}

}
