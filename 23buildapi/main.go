package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// models for courses -file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// (Whenever we make a fake database and make a databse that called a seeding):
// seeding a pseudo-random number generator (PRNG) or seeding a database
// fake db
var courses []Course

// middleware, helper - file
func (c *Course) ISEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("We Will Build Our First APi!!-MyGolang")

	// routing part
	r := mux.NewRouter()

	//seeding part
	courses = append(courses, Course{CourseId: "1", CourseName: "ReactJS", CoursePrice: 899, Author: &Author{Fullname: "Akash Kalsariya", Website: "React.dev"}})

	courses = append(courses, Course{CourseId: "2", CourseName: "NodeJS", CoursePrice: 899, Author: &Author{Fullname: "Akash Kalsariya", Website: "Node.dev"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("Get")
	r.HandleFunc("/courses", getallcourses).Methods("Get")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("Get")
	r.HandleFunc("/course", craeteOneCourse).Methods("Post")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("Put")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("Delete")

	//Listen  to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - file

//serve Home Route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome  to api by ak07</h1>"))
}

func getallcourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses!!")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through courses,find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found With Given Id")
	return
}

func craeteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	//what if:body is Empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please SEnd Some Data!! ")
	}

	//what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.ISEmpty() {
		json.NewEncoder(w).Encode("No Data inside Json!!")
		return
	}

	//todo:  If title is duplicate
	//loop,title matches with coursename,Json
	for _, existingCourse := range courses {
		fmt.Println("Comparing with:", existingCourse.CourseName)

		if strings.EqualFold(existingCourse.CourseName, course.CourseName) {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Course already exists: " + course.CourseName,
			})
			return
		}
	}

	//gennrate a unique id ,string
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One Course!!")
	w.Header().Set("Content-type", "application/json")

	//first - grab id from req
	params := mux.Vars(r)

	//loop, id , remove, add with my Id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
		//todo : send a response when id is not found
		// ✅ TODO COMPLETED: send response when id not found
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Course not found",
		})
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Course!!")
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {

			// delete course
			courses = append(courses[:index], courses[index+1:]...)

			// ✅ send response
			// ✅ send JSON response
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Course is deleted",
				"id":      params["id"],
			})
			return
		}
	}

	// ✅ if not found
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "Course not found",
	})
}
