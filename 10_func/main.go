package main

import "fmt"

func main() {
	greeter()
	fmt.Println(adder(1, 4))
	fmt.Println(proAdder(3, 5, 3, 6))
}

func greeter() {
	fmt.Println("Hello Everyone")
}

func adder(valueOne int, valueTwo int) int {
	return valueOne + valueTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "This is the sum of all the values"
}
