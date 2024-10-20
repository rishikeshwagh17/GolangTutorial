package main

import (
	"fmt"
	"net/url"
)

// URL
const myUrl string = "https://lco.dev:3000/learn/learn?coursename=reactjs&paymentId=adadaaa"

func main() {

	fmt.Println("Urls baisc in golang")
	fmt.Println(myUrl)

	//parse the url
	result, _ := url.Parse(myUrl)
	fmt.Printf("Result type %T \n", result)
	fmt.Println("result", result.Scheme)
	fmt.Println("result", result.Host)
	fmt.Println("result", result.Path)
	fmt.Println("result", result.Port())
	fmt.Println("result", result.RawQuery)

	//store queryParam
	queryParam := result.Query()
	// now loop over all the queryParams
	for _, val := range queryParam {
		fmt.Println("param is ", val)
	}

	//if we have all info  and we want to create URl from it
	constructUrl := &url.URL{
		Scheme:   "https",
		Path:     "/learn",
		Host:     "lco.dev",
		RawQuery: "user=rishi",
	}

	anotherUrl := constructUrl.String()
	fmt.Println("new Url ", anotherUrl)
}
