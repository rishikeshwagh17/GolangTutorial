package main

import (
	"fmt"
)

func main()  {
	fmt.Println("hello to array in golang")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Pomogrenate"
	
	fmt.Println("Fruit list is: ",fruitList)
	fmt.Println("fruit list length is: ",len(fruitList))
	fmt.Println("value of 3rd fruit is: ",fruitList[2])

	veggiesList := [4]string{"potato","mushroom","loki","ladyFinger"}

	fmt.Println("veggies list",veggiesList)
	fmt.Println("veggies list length is ",len(veggiesList))

	// var ptr *int;
	// num := 32

	// ptr = &num
	// fmt.Println("value of pointer is",ptr)
	// fmt.Println("value of pointer is",*ptr)
}