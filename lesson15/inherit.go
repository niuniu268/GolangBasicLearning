package main

import "fmt"

func main() {

	a := Human{
		name: "xiaoA",
		sex:  "male",
	}

	a.Eat()
	a.work()

	fmt.Println("=========")

	b := Superman{
		Human: Human{
			name: "xiaoB",
			sex:  "female",
		},
		level: 0,
	}

	b.work()
	b.Eat()
	b.Fly()

	var c Superman

	c.name = "xiaoC"
	c.sex = "none"
	c.level = 2

	c.Eat()
	c.work()
	c.Fly()

}
