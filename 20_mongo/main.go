package main

import (
	"fmt"
	"log"
	"net/http"
	"netflix_api_mock/router"
)

func main() {
	fmt.Println("MONGO DB API STARTING....")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server up and running at 8080")
}
