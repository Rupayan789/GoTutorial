package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string   `json:"coursename"`
	Password string   `json:"-"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	courses := []Course{
		{
			"ReactJSBootcamp", "123456", 399, "udemy", []string{"Web Development", "Frontend Development"},
		}, {
			"GoBootcamp", "234567", 599, "coursera", []string{"Backend Development", "Full-stack Development"},
		}, {
			"PythonBootcamp", "345678", 499, "udemy", nil,
		},
	}
	EncodeJson(courses)

}

func EncodeJson(data []Course) {
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(jsonData))

}
