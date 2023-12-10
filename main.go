package main

import "fmt"

type person struct {
	name string
	family string
	age int
}

func main() {
	var p person

	p.age = 30
	p.family = "babayy"
	p.name = "mohammadjavad"

	fmt.Println("age : ", p.age, "name : ",p.name, "family : ",p.family)
}
