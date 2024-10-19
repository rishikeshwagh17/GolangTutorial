package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to the maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "javascript"
	languages["RB"] = "ruby"
	languages["PY"] = "python"

	fmt.Println("List of all languages",languages)
	fmt.Println("JS shorts for ",languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)

	//loops in golang 
	for _,value := range languages {
		fmt.Printf("for the key value is %v \n",value)
	}
}