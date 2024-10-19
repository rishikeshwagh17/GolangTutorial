package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello to the learning of slices in golang")
	fruitList := []string{"Apple","Tomato","Peach"}  //this is slice means array like but not array because we have to give length each time for array

	fruitList = append(fruitList, "Mango","Banana")
	// fmt.Println("my fruitList is ",fruitList)
	// fmt.Printf("Type of fruitList is %T \n",fruitList)

	//to get particular portion or slice from slice we use below syntax
	// fmt.Println(append(fruitList[:3]))
	// fmt.Println(append(fruitList[1:3]))
	// fmt.Println(append(fruitList[1:]))

	highScores := make([]int,3)
	highScores[0] = 10
	highScores[1] = 45
	highScores[2] = 90
	// highScores[3] = 123 // this will not work

	highScores = append(highScores, 444,565,123)
	// fmt.Println("highscores slice ",highScores)

	sort.Ints(highScores)
	// fmt.Println(highScores)

	courses := make([]string, 0)
	courses = append(courses, "reactsjs","nodejs","swift","python","ruby")
	//now we want to remove element at index 2 from slice way to remove it is we have to use append only

	var index = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("courses value after removal of 2nd element ",courses)
}