// GO server
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome to web")
	PerformGetRequest()
	PerformPostJsonRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:4000/api/hello"
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status code", response.StatusCode)
	fmt.Println("content length is ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is ", byteCount)

	fmt.Println("response from api is ", responseString.String())
}

func PerformPostJsonRequest() {
	const myUrl = "http://localhost:4000/api/data"

	//fake json payload
	requestBody := strings.NewReader(`
		{	
			"name" : "Rishi",
			"age" : 12
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is ", byteCount)

	fmt.Println("response from api is ", responseString.String())
}
