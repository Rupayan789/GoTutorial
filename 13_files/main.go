package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go into a file named hello.txt"

	file, err := os.Create("./hello.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Length is ", length)

	defer file.Close()

	readFile("./hello.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Date from the file :", string(databyte))
}
