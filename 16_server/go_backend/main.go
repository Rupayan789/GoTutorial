package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Working with web server")
	// Initialize server

	var webserver WebServer = WebServer{"http://localhost:3000"}

	// webserver.PerformSimpleGetRequest()
	// webserver.PerformPostWithFakeJsonRequest()
	webserver.PerformPostWithFormDataRequest()

}

type WebServer struct {
	Baseurl string
}

func (ws WebServer) PerformSimpleGetRequest() {
	response, _ := http.Get(ws.Baseurl)

	dataBytes, _ := io.ReadAll(response.Body)

	defer response.Body.Close()

	// Working with String Builder

	var responseString strings.Builder

	byteCount, _ := responseString.Write(dataBytes)

	fmt.Println("Bytes read =", byteCount)

	responseContent := responseString.String()

	fmt.Println("Response Content =", responseContent)
}

func (ws WebServer) PerformPostWithFakeJsonRequest() {
	payload := strings.NewReader(`
		{
			"id":4,
			"name":"Joe",
			"age":23
		}
	`)

	response, err := http.Post(ws.Baseurl, "application/json", payload)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	dataBytes, _ := io.ReadAll(response.Body)
	var responseString strings.Builder
	_, _ = responseString.Write(dataBytes)

	content := responseString.String()

	fmt.Println("Response of POST\n", content)

}

func (ws WebServer) PerformPostWithFormDataRequest() {
	payload := url.Values{}

	payload.Add("id", "5")
	payload.Add("name", "Kane")
	payload.Add("age", "35")

	response, err := http.PostForm(ws.Baseurl, payload)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	dataBytes, _ := io.ReadAll(response.Body)
	var responseString strings.Builder
	_, _ = responseString.Write(dataBytes)

	content := responseString.String()

	fmt.Println("Response of POST\n", content)
}
