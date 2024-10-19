package main

import "fmt"

func main() {
	fmt.Println("if else in go lang")
	loginCount := 23

	var result string
	if loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 && loginCount < 30 {
		result = "Member"
	} else {
		result = "Owner"
	}

	fmt.Println("the user is ", result)

	for num := 10; num <= 20; {
		if num == 10 {
			fmt.Print(num)
		} else {
			fmt.Print(" ", num)
		}
		num++
	}
}
