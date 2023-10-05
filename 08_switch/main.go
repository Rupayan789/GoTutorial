package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// Working with switch
	var diceRoll = rand.Intn(6) + 1
	switch diceRoll {
	case 1:
		fmt.Println("You can move 1 step")
	case 2:
		fmt.Println("You can move 2 steps")
	case 3:
		fmt.Println("You can move 3 steps")
		fallthrough
	case 4:
		fmt.Println("You can move 4 steps")
		fallthrough
	case 5:
		fmt.Println("You can move 5 steps")
	case 6:
		fmt.Println("You can move 6 steps and Try again!")
	default:
		fmt.Println("What was that!")
	}

}
