package main

import "fmt"

func main() {
	daysList := []string{"Sun", "Mon", "Tue", "Wed", "Thurs", "Fri", "Sat"}

	for i := 0; i < len(daysList); i++ {
		fmt.Println(daysList[i])
	}

	fmt.Println()
	colorList := []string{"Red", "Green", "Yellow"}
	for i := range colorList {
		fmt.Println(colorList[i])
	}

	fmt.Println()
	monthsList := []string{"Jan", "Feb", "Mar"}
	for _, value := range monthsList {
		fmt.Println(value)
	}

	// break and continue

	i := 1
	for i < 10 {
		if i == 7 {
			break
		} else if i == 3 {
			i++
			continue
		} else if i == 5 {
			goto label
		}
		fmt.Println(i)
		i++
	}

label:
	fmt.Println("Goto reached here")

}
