package main

import "fmt"

//Package scope array definition
var integerSlice []int

var stringSlice []string

func main() {

	integerSlice = []int{10, 20, 30, 40}

	fmt.Println("This is the integer Slice: ", integerSlice)

	stringSlice = []string{"first", "second", "third"}

	fmt.Println("This is the string slice: ", stringSlice)

	printInteger()

	s := []int{10, 20, 30, 40}
	fmt.Println("len(s) :", len(s), "cap(s) : ", cap(s))
	fmt.Println(s[:])


	for key, value := range s {
	fmt.Println(key, value)
	}

	fmt.Println("======================")

	for i := 0; 1< len(s); i++ {
		fmt.Println(i,s[i])
	}
}

//This function can access the integerSlice because integerSlice is package scoped
func printInteger() {
	fmt.Println("Print the integers: ", integerSlice)
}
