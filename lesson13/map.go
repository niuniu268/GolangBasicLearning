package main

import "fmt"

func main() {
	// declaration 1
	map0 := make(map[int]string, 0)

	map0[1] = "test1"
	map0[2] = "test2"
	map0[3] = "test3"

	// declaration 2
	var map1 map[int]string

	map1 = make(map[int]string, 4)

	map1[1] = "test1"
	map1[2] = "test2"
	map1[3] = "test3"

	// declaration 3

	map2 := map[int]string{
		1: "test1",
		2: "test2",
		3: "test3",
	}

	fmt.Println(map2)

	for index, value := range map2 {

		fmt.Println(index)
		fmt.Println(value)

	}

}
