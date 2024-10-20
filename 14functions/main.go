package main

import "fmt"

func main() {
	fmt.Println("functions in golang")
	greeter()
	addition := adder(2, 5)
	fmt.Println(addition)

	add := proAdder(1, 2, 4, 77, 88, 32, 11)
	fmt.Println("addition is ", add)

	msg1, msg2 := message()
	fmt.Println("the value of message 1 is ", msg1)
	fmt.Println("the value of message 2 is ", msg2)
}

func greeter() {
	fmt.Println("Hello from Rishi learning")
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

//now what we have more paramteres in function then we can use
func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

//function can return more than 1 value in golang
func message() (string, string) {
	return "hello", "Welcome"
}
