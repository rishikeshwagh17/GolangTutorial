package main

import "fmt"

func main() {
	// fmt.Println("defer keyword in golang")
	// defer fmt.Println("Hello world")  this code will work as normal

	// defer fmt.Println("Hello world") // defer put this code exactly above return or end of the function definition
	// fmt.Println("defer in golang")

	//in this case it will act like a stack Last in first out so last defer will execute first
	defer fmt.Println("world")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello") // op- hello Two One world

	//what we have onr function call having defer keyword
	//here main start executing code line by line Hello will print then myDefer() is not defer execution so it is next and then all with defer stack

	// o/p - Hello (myDefer - 3210) Two One world
	myDefer()
}

func myDefer() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}
