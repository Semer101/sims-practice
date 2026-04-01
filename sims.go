package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

type Student struct {
	ID string `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Age int `json:"age"`
}

var students []Student

func GetStudent(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, student := range students {
		if student.ID == params["id"] {
			json.NewEncoder(w).Encode(student)
			return
		}
	}
	json.NewEncoder(w).Encode(&Student{})
}

func GetStudents(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(students)
}

func CreateStudent(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var student Student
	_ = json.NewDecoder(req.Body).Decode(&student)
	student.ID = params["id"]
	students = append(students, student)
	json.NewEncoder(w).Encode(students)
}

func DeleteStudent(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, student := range students {
		if student.ID == params["id"] {
			students = append(students[:index], students[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(students)
}

func UpdateStudent(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var updatedStudent Student
	_ = json.NewDecoder(req.Body).Decode(&updatedStudent)
	for index, student := range students {
		if student.ID == params["id"] {
			updatedStudent.ID = student.ID
			students[index] = updatedStudent
			json.NewEncoder(w).Encode(updatedStudent)
			return
		}
	}
	http.Error(w, "Student not found", http.StatusNotFound)
}

func PatchStudent(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var updatedStudent Student
	_ = json.NewDecoder(req.Body).Decode(&updatedStudent)
	for index, student := range students {
		if student.ID == params["id"] {
			if updatedStudent.FirstName != "" {
				student.FirstName = updatedStudent.FirstName
			}
			if updatedStudent.LastName != "" {
				student.LastName = updatedStudent.LastName
			}
			if updatedStudent.Age != 0 {
				student.Age = updatedStudent.Age
			}
			students[index] = student
			json.NewEncoder(w).Encode(student)
			return
		}
	}
	http.Error(w, "Student not found", http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()

	students = append(students, Student{ID: "1", FirstName:"Sabrin", LastName:"Abdul", Age:22})
	students = append(students, Student{ID: "2", FirstName:"Sara", LastName:"Jemal", Age:21})
	students = append(students, Student{ID: "3", FirstName:"Ahmed", LastName:"Omar", Age:23})

	router.HandleFunc("/students", GetStudents).Methods("GET")
	router.HandleFunc("/students/{id}", GetStudent).Methods("GET")
	router.HandleFunc("/students/{id}", CreateStudent).Methods("POST")
	router.HandleFunc("/students/{id}", DeleteStudent).Methods("DELETE")
	router.HandleFunc("/students/{id}", UpdateStudent).Methods("PUT")
	router.HandleFunc("/students/{id}", PatchStudent).Methods("PATCH")

	http.ListenAndServe(":8080", router)
}