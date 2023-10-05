package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=327283orb"

func main() {
	fmt.Println("Welcome to handling url in Golang")
	result, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Url result is of type %T and value is %+v\n", result, result)

	fmt.Println(result.Host)
	fmt.Println(result.Scheme)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Printf("Type of Query Params is %T\n", qparams)

	for key, value := range qparams {
		fmt.Printf("Key is %v and value is %v\n", key, value[0])
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)

}
