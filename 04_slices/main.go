package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Working with slices

	var fruitList = []string{"Apple", "Orange", "Avocado"}
	sort.Strings(fruitList)
	fmt.Println(fruitList)
	var newFruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(newFruitList, fruitList)

	fmt.Println(newFruitList[1:4])

	// Dyanmic Allocation

	var carList = make([]string, 5)
	carList[0] = "Mercedes"
	carList[1] = "Audi"
	carList[2] = "Suzuki"
	carList[3] = "Ferrari"
	carList[4] = "Porsche"
	sort.Strings(carList)
	fmt.Println(carList, len(carList))
	var myCars string = strings.Join(carList, "-")
	fmt.Println(myCars)

	// How to remove an item based on index
	var index int = 2
	fmt.Println(append(carList[:index], carList[index+1:]...))

}
