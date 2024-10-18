package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to conversion tutorial")

	fmt.Println("welcome to our Pizza app")
	fmt.Println("please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for rating, ",input)

	numRating,err := strconv.ParseFloat(strings.TrimSpace(input),32)
	
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println("Added 1 to your rating: ",numRating + 1)
	}

}