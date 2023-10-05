package main

import "fmt"

func main() {
	loginCount := 10
	var result string
	// Only this syntax works,wont work if brackets are below if
	if loginCount < 10 {
		result = "Regular"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if i := 5; i%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

}
