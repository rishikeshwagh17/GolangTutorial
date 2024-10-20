package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in golang")

	content := "This is the text related to file working in golang"

	file, err := os.Create("./myFile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkHasError(err)
	len, err := io.WriteString(file, content)
	checkHasError(err)
	fmt.Println("lngth of content inside files is ", len)
	defer file.Close()
	readFile("./myFile.txt")
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkHasError(err)
	fmt.Println("Text data inside the file is \n", string(dataByte))
}

// common way to use error function
func checkHasError(err error) {
	if err != nil {
		panic(err)
	}
}
