package main

import "fmt"

type User struct {
	Name   string
	email  string
	Age    int
	Status bool
}

func main() {
	var user User = User{"Rupayan", "abc@gmail.com", 22, true}
	fmt.Println(user)
	fmt.Printf("The user is %+v", user)
}
