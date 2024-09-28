package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for courses:
type Course struct {
	CourseId   string  `json:"courseid"`
	CourseName string  `json:"coursename"`
	CoursePrice int    `json:"price"`
	Author     *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helpers
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

// controllers - file

// server home:
func ServerHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello there!</h1>"))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses req called")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func GetSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get single course req called")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// looping through courses:
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	http.Error(w, "No course found with given id!", http.StatusNotFound)
}

func CreateNewCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create new course req called")
	w.Header().Set("Content-Type", "application/json")

	// case if body is empty
	if r.Body == nil {
		http.Error(w, "No data found", http.StatusBadRequest)
		return
	}

	var course Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	if course.IsEmpty() {
		http.Error(w, "No data found", http.StatusBadRequest)
		return
	}

	// generate Unique id, string
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(course)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating a particular course!")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
				http.Error(w, "Invalid data", http.StatusBadRequest)
				return
			}

			course.CourseId = params["id"]
			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
			return
		}
	}

	http.Error(w, "No course with the given Id", http.StatusNotFound)
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete course req called")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course Deleted Successfully")
			return
		}
	}

	http.Error(w, "No course with given id found!", http.StatusNotFound)
}

func main() {
	fmt.Println("Server Started on 4000")
	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "1", CourseName: "Reactjs", CoursePrice: 299, Author: &Author{FullName: "Anurag Kamboj", Website: "https://kamboj.me"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Mern", CoursePrice: 499, Author: &Author{FullName: "Anurag Kamboj", Website: "https://kamboj.me"}})

	r.HandleFunc("/", ServerHome).Methods("GET")
	r.HandleFunc("/getall", GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", GetSingleCourse).Methods("GET")
	r.HandleFunc("/course", CreateNewCourse).Methods("POST")
	r.HandleFunc("/course/{id}", UpdateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", DeleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}
