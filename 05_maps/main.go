package main

import "fmt"

func main() {
	// Working with Maps

	var languages = make(map[string]string)
	languages["JS"] = "javascript"
	languages["RB"] = "ruby"
	languages["JV"] = "java"
	languages["PY"] = "python"
	fmt.Println(languages)
	delete(languages, "JV")
	fmt.Println(languages)
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}
}
