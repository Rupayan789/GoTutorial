package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", removeOneCourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))

}

type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"coursePrice"`
	Author      *Author `json:"author"`
}

type Author struct {
	AuthorId string `json:"authorId"`
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

func (c *Course) isEmpty() bool {
	return c.CourseName == ""
}

func (a *Author) isEmpty() bool {
	return a.FullName == ""
}

func (c *Course) getField(field string) reflect.Value {
	r := reflect.ValueOf(c)
	f := reflect.Indirect(r).FieldByName(field)
	fmt.Printf("%v and %T", f, f)
	return f
}

var authors []Author = []Author{
	{"1", "John", "john.dev"},
	{"2", "Joe", "joe.dev"},
	{"3", "Henry", "henry.dev"},
}

var courses []Course = []Course{
	{
		"1", "ReactJs Bootcamp", 599, &authors[0],
	},
}

// var courses []Course = []Course{}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
	return
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting one course")
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating one course")
	w.Header().Set("Content-Type", "application/json")
	content := r.Body
	if content == nil {
		json.NewEncoder(w).Encode("No data received")
		return
	}
	var course Course
	_ = json.NewDecoder(content).Decode(&course)

	if course.isEmpty() {
		json.NewEncoder(w).Encode("Invalid course data received")
		return
	}

	// create id as number and convert it to string
	// push course to courses

	course.CourseId = strconv.Itoa(rand.Intn(10000))

	courses = append(courses, course)

	json.NewEncoder(w).Encode("Succesfully created course")
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating one course")
	params := mux.Vars(r)
	body := r.Body
	// var updateKeyValues map
	var updatedCourse Course
	json.NewDecoder(body).Decode(&updatedCourse)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			updatedCourse.CourseId = course.CourseId
			courses = append(courses, updatedCourse)
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}
	json.NewEncoder(w).Encode("Course id invalied")
}

func removeOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Removing one course")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Successfully removed")
			return
		}
	}
	json.NewEncoder(w).Encode("Didn't find one!")
}
