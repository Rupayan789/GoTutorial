package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const url string = "https://lco.dev"

func main() {
	fmt.Println("Making Web Requests")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type %T\n", response)

	defer response.Body.Close() // callers responsibility to close the connection
	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(dataBytes)

	var filename string = "./lco.html"

	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("Content written to lco.txt of length:", length)
	defer file.Close()
}
