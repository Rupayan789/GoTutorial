package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("This module covers User Input")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please give your rating")
	rating, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us ", rating)
	new_rating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)
	if err != nil {
		fmt.Println("Error occured:", err)
	}
	fmt.Println("New rating is ", new_rating+1)

}
