package main

import "fmt"

type person struct {
	name string
	family string
	age int
}

func main() {
	var p person = person{name: "mohammadjavad", family: "babayy", age: 20}

	fmt.Println("age : ", p.age, "name : ",p.name, "family : ",p.family)
}
