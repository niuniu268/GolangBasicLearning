package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id  int    `json:"id,omitempty" doc:"another tag"`
	Sex string `json:"se,omitempty"`
}

func findTag(str interface{}) {

	T := reflect.TypeOf(str)
	for i := 0; i < T.NumField(); i++ {

		T.Field(i).Tag.Get("json")
		fmt.Println(T.Field(i).Tag.Get("json"))

	}

}

func main() {

	test := User{
		Id:  2,
		Sex: "male",
	}

	findTag(test)

}
