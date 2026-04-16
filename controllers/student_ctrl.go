package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"sims-practice/database"
	"sims-practice/models"

	"github.com/gorilla/mux"
)

var templates *template.Template

func LoadTemplates() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

func Home(w http.ResponseWriter, r *http.Request) {
    var students []models.Student
    database.DB.Find(&students)

    // Wrap the data in a map so the template can find ".Students"
    data := map[string]interface{}{
        "Students": students,
    }

    if r.Header.Get("HX-Request") == "true" {
        // Only send the part inside the define block
        templates.ExecuteTemplate(w, "students-page", data)
    } else {
        // Send the whole page for the first visit
        templates.ExecuteTemplate(w, "index.html", data)
    }
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	age, _ := strconv.Atoi(r.FormValue("age"))
	student := models.Student{
		FirstName: r.FormValue("firstname"),
		LastName:  r.FormValue("lastname"),
		Age:       age,
	}

	models.SanitizeStudent(&student)
	if err := models.Validate.Struct(student); err != nil {
		http.Error(w, "Invalid Input", 400)
		return
	}

	database.DB.Create(&student)
	templates.ExecuteTemplate(w, "student-row", student)
}

func EditStudentForm(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var student models.Student
	database.DB.First(&student, id)
	templates.ExecuteTemplate(w, "student-edit-row", student)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var student models.Student
	database.DB.First(&student, id)

	student.FirstName = r.FormValue("firstname")
	student.LastName = r.FormValue("lastname")
	student.Age, _ = strconv.Atoi(r.FormValue("age"))

	models.SanitizeStudent(&student)
	database.DB.Save(&student)
	templates.ExecuteTemplate(w, "student-row", student)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	database.DB.Delete(&models.Student{}, id)
	w.WriteHeader(http.StatusOK)
}
