package main

import (
	"fmt"
	"log"
	"net/http"

	"sims-practice/controllers"
	"sims-practice/database"
	"sims-practice/models"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	controllers.LoadTemplates()
}

func main() {
	database.Connect()

	// Migrate models
	database.DB.AutoMigrate(&models.Student{}, &models.Teacher{}, &models.Grade{})

	r := mux.NewRouter()

	// Logging Middleware
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	})

	// Routes
	r.HandleFunc("/", controllers.Home).Methods("GET")
	r.HandleFunc("/students", controllers.CreateStudent).Methods("POST")
	r.HandleFunc("/students/edit/{id}", controllers.EditStudentForm).Methods("GET")
	r.HandleFunc("/students/update/{id}", controllers.UpdateStudent).Methods("PUT")
	r.HandleFunc("/api/students/{id}", controllers.DeleteStudent).Methods("DELETE")

	// --- Teachers ---
	r.HandleFunc("/teachers", controllers.ListTeachers).Methods("GET")
	r.HandleFunc("/teachers", controllers.CreateTeacher).Methods("POST") // Removed /add to match HTML
	r.HandleFunc("/api/teachers/{id}", controllers.DeleteTeacher).Methods("DELETE") // Added /api/

	// --- Grades ---
	r.HandleFunc("/grades", controllers.ListGrades).Methods("GET")
	r.HandleFunc("/grades/assign", controllers.AssignGrade).Methods("POST")
	r.HandleFunc("/api/grades/{id}", controllers.DeleteGrade).Methods("DELETE") // Added /api/

	fmt.Println("Server active at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}