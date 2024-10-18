package main

import "fmt"

func main() {
	fmt.Println("wlcome to the class of pointers")

	// var ptr *int
	// fmt.Println("value of pointer is ",ptr)

	// myNumber := 23
	// var ptr = &myNumber

	// fmt.Println("value of actual pointer is ",ptr)
	// fmt.Println("value of actual pointer is", *ptr)

	var ptr *int
	myNumber := 43
	ptr = &myNumber

	fmt.Println("value of pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("value of myNumber now is ",myNumber)

}