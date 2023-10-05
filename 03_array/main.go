package main

import (
	"fmt"
)

func main() {
	// This module also covers Pointers and time values

	// Working with Array

	var fruit_list = [4]string{"Apple", "Orange", "Mango"}
	fmt.Println(fruit_list)

	var animal_list [3]string
	animal_list[0] = "Cat"
	animal_list[2] = "Dog"
	fmt.Println(animal_list)

	// Pointers
	var number int64 = 10
	var ptr *int64 = &number
	var ptr2 *int
	fmt.Println(ptr2, ptr)

}
