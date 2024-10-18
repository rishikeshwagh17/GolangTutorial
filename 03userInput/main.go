package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello welcome to go tutorial user input"
	fmt.Println(welcome)

	//rading from input source

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza : " )

	//comma ok || err err  just like try catch in other language
	input, _  := reader.ReadString('\n')

	fmt.Println("thanks for rating, ",input)
	fmt.Printf("Type of this rating is %T ", input)
}