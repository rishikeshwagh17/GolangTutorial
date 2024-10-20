package main

import "fmt"

func main() {
	fmt.Println("methods in golang")
	myUser := User{"rishi", "r123@mail.com", true, 23}
	myUser.MyStructMethod()
	// myUser.NewMail()
	myUser.NewMailNew()
	fmt.Println("modified User data is", myUser)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

//Always use Capital letter for method
func (u User) MyStructMethod() { //this is method as it hold reference to our struct and we can call it similar to class method in java as no class is golang
	fmt.Println("status of the user is ", u.Status)
	fmt.Println("name of the user is ", u.Name)
	fmt.Println("age of the user is ", u.Age)
}

//we can also manipulate the struct data
func (u User) NewMail() {
	u.Email = "RishiNewMail@.com"
	fmt.Println("Email of the user is ", u.Email)
} //but this method is not manipulating the original myUser struct beecause we this we are pass copy of user struct and the actual reference

//for that we need pointers and we need to pass pointer to function
func (u *User) NewMailNew() {
	u.Email = "RishiNewMail@.com"
	fmt.Println("Email of the user is ", u.Email)
} //but this method is not manipulating the original myUser struct beecause we this we are pass copy of user struct and the actual reference
