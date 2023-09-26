package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	Price int    `json:"price"`
}

func main() {

	movie := Movie{
		Title: "a film",
		Year:  2000,
		Price: 100,
	}

	var mymovie Movie

	//	convert to json

	marshal, error := json.Marshal(movie)

	if error != nil {
		fmt.Println("error", error)
		return
	} else {
		fmt.Printf("%s\n", marshal)
	}

	json.Unmarshal(marshal, &mymovie)

	fmt.Printf("%v\n", mymovie)

}
