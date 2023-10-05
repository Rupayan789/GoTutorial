package main

import "fmt"

func main() {
	fmt.Println(greeter())
	count()
	fmt.Println("Count is called")
}

func greeter() string {
	defer fmt.Println("World")
	fmt.Println("Hello")
	return "Bye"
}
func count() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
