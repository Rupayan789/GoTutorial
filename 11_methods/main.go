package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("Status of the user is", u.Status)
}

func (u User) SendEmail() {
	u.Email = "test@email.com"
	fmt.Printf("%+v\n", u)
}

func main() {
	var user User = User{"Rupayan", "abc@gmail.com", 22, true}
	fmt.Println(user)
	user.GetStatus()
	user.SendEmail()
	fmt.Printf("The user is %+v", user)

}
