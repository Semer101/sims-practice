package controllers

import (
	"net/http"
	"sims-practice/database"
	"sims-practice/models"
	"github.com/gorilla/mux"
)

func CreateTeacher(w http.ResponseWriter, r *http.Request) {
	teacher := models.Teacher{
		Name:       r.FormValue("name"),
		Department: r.FormValue("department"),
	}
	
	if err := database.DB.Create(&teacher).Error; err != nil {
		http.Error(w, "Database error", 500)
		return
	}
	
	// Returns just the <tr> so HTMX hx-swap="beforeend" works
	templates.ExecuteTemplate(w, "teacher-row", teacher)
}

func ListTeachers(w http.ResponseWriter, r *http.Request) {
	var teachers []models.Teacher
	database.DB.Find(&teachers)

	data := map[string]interface{}{
		"Teachers": teachers,
	}

	templates.ExecuteTemplate(w, "teachers-page", data)
}

func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	database.DB.Delete(&models.Teacher{}, id)
	w.WriteHeader(http.StatusOK)
}