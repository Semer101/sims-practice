package controllers

import (
	"net/http"
	"strconv"

	"sims-practice/database"
	"sims-practice/models"

	"github.com/gorilla/mux"
)

func AssignGrade(w http.ResponseWriter, r *http.Request) {
	sID := parseUint(r.FormValue("student_id"))
	tID := parseUint(r.FormValue("teacher_id"))

	grade := models.Grade{
		Value:     r.FormValue("grade"),
		StudentID: sID,
		TeacherID: tID,
	}

	if err := database.DB.Create(&grade).Error; err != nil {
		http.Error(w, "Could not assign grade. Ensure Student and Teacher exist.", 500)
		return
	}

	// Fetch the record again with Preload so we have the Student/Teacher names for the UI
	database.DB.Preload("Student").Preload("Teacher").First(&grade, grade.ID)

	// Return the row fragment to the HTMX target
	templates.ExecuteTemplate(w, "grade-row", grade)
}

func ListGrades(w http.ResponseWriter, r *http.Request) {
	var students []models.Student
	var teachers []models.Teacher
	var grades []models.Grade

	database.DB.Find(&students)
	database.DB.Find(&teachers)
	database.DB.Preload("Student").Preload("Teacher").Find(&grades)

	data := map[string]interface{}{
		"Students": students,
		"Teachers": teachers,
		"Grades":   grades,
	}

	templates.ExecuteTemplate(w, "grades-page", data)
}

func DeleteGrade(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	database.DB.Delete(&models.Grade{}, id)
	w.WriteHeader(http.StatusOK)
}

func parseUint(s string) uint {
	val, _ := strconv.ParseUint(s, 10, 32)
	return uint(val)
}
