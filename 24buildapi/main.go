package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//  Model for courses:
type Course struct{
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}


type Author struct{
	FullName string `json:"fullname"`
	Website string `json:"website"`
}



// fake DB
var courses [] Course

// middleware, helpers;
func (c* Course) IsEmpty() bool{
 return (c.CourseName =="")
}

// controllers  - file

// server home:
func ServerHome(w http.ResponseWriter,r *http.Request){
   w.Write([]byte ("<h1>Hello there!</h1>"))
}


func GetAllCourses(w http.ResponseWriter,r *http.Request){
	fmt.Println("get all courses req called")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}


func GetSingleCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("get single course req called")
	w.Header().Set("Content-Type","application/json")

    params := mux.Vars(r)

	// looping through courses:
	for _,course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id!")
}


func CreateNewCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("get single course req called")
	w.Header().Set("Content-Type","application/json")

    // case if  body  is empty()
	if r.Body == nil{
		json.NewEncoder(w).Encode("No data found")
		return
	}

	var course Course
	_= json.NewDecoder(r.Body).Decode(&course)

	if(course.IsEmpty()){
		json.NewEncoder(w).Encode("No data found")
		return
	}


	// generate Unique id, string

	rand.New(rand.NewSource(time.Now().UnixNano()))
	course.CourseId = strconv.Itoa(rand.Intn(100))


	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}


func Updatecourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Updating a particular course!")
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	for index, course  := range courses{
		if course.CourseId== params["id"]{
			courses = append(courses[:index],courses[index+1:]...)
			var course Course
			_= json.NewDecoder(r.Body).Decode(&course)

			course.CourseId = params["id"]
			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course with the given Id")

}


func DeleteCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("get single course req called")
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)

	for index,course := range courses{
       if(course.CourseId == params["id"]){
		courses = append(courses[:index],courses[index+1:]...)
		json.NewEncoder(w).Encode("Course Deleted Successfully")
		return
	   }
	}

	json.NewEncoder(w).Encode("No course with given id found!")
}


func main()  {
	
}