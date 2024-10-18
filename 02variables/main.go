package main

import "fmt"

const LoginToken string = "ABBCFEEERRRA" // this is public var and can be used any file inside this folder
func main() {
	var username string = "Rishi"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T \n",username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variables is of type : %T \n ",isLoggedIn)
	
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variables is of type : %T \n ", smallVal)

	var smallFloat float64 = 255.55453434312231213
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type :  %T \n ",smallFloat)


	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)  //it doesnt store garabge value when we are not passing any value while declaring a var

	//implicit type mean we dont need to tell type of variable lexer automatically take it
	var website = "www.google.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type : %T \n", website)
	
	//no var type - no var leyword
	//this := is called short variable declaration operator & it is only work inside function outisde function definition you need var keyword
	numberOfUsers := 300000.00
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)

}