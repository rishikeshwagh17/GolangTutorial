package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://catfact.ninja/fact"

func main() {
	//now do a http request
	//for that we have http package

	response, err := http.Get(url)

	//handle err
	if err != nil {
		panic(err)
	}
	//else check type of response and work with response
	fmt.Printf("the type of response is %T \n", response)
	defer response.Body.Close()
	//now use response here we simply use ioUtil package
	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("response content ", string(dataBytes))

}
