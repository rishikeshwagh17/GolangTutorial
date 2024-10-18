package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Hello we are learning about time")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2024,time.August,17,23,23,0,0,time.Local)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))
}