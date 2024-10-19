package main

import "fmt"

func main() {
	fmt.Println("Hello to learning of structs in golang")
	//no inheritance in golang no super or parent
	rishi := User{"Rishi", "rishikesh123@gmail.com", true, 25}
	fmt.Printf("Rishikesh struct details are %+v \n", rishi)

	//you want to access particular fields then use
	fmt.Printf("Name is %v and email is %v ", rishi.Name, rishi.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
