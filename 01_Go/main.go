package main

import "fmt"

// mod := 122423423 // does not works

func main() {
	fmt.Println("Hello I am from another world")
	var username string = "Rupayan789"
	fmt.Printf("%s is of type %T \n", username, username)
	var phone_no int64 = 8249742079
	fmt.Printf("%d is of type %T \n", phone_no, phone_no)
	var is_logged_in bool = true
	fmt.Printf("%t is of type %T \n", is_logged_in, is_logged_in)
	var balance float64 = 100.23434
	fmt.Printf("%f is of type %T \n", balance, balance)

	pincode := 755019
	fmt.Println(pincode)

	var any_value int8
	fmt.Println(any_value) // default value is 0 unless specified
}
