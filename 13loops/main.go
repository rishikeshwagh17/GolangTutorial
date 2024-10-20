package main

import "fmt"

func main() {
	fmt.Println("welcome to the loops in golang")

	// for i := 0; i < 10; i++ {
	// 	fmt.Print(i, " ")

	// }
	// fmt.Println("")
	// days := []string{"sunday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

	// for i := range days { //here i is index from range
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days { //day here is value from range and _ is index it is one of syntax
	// 	fmt.Printf(" index is %v and value is %v \n", index, day)
	// }
	rougeValue := 1

	for rougeValue < 10 {
		if rougeValue == 9 {
			goto myGoto
		}
		if rougeValue == 5 {
			rougeValue++
			continue
		}

		fmt.Println("value is ", rougeValue)
		rougeValue++
	}
myGoto:
	fmt.Println("this is my go to condition in loop")
}
