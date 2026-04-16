package models // Changed from main to models

import (
    "github.com/go-playground/validator/v10"
    "github.com/microcosm-cc/bluemonday"
)

type Student struct {
    ID        uint   `gorm:"primaryKey" json:"id"`
    FirstName string `json:"firstname" validate:"required,min=2"`
    LastName  string `json:"lastname" validate:"required,min=2"`
    Age       int    `json:"age" validate:"gte=0,lte=120"`
}

var Validate = validator.New() // Capitalized to export it

func SanitizeStudent(s *Student) { // Capitalized to export it
    p := bluemonday.StrictPolicy()
    s.FirstName = p.Sanitize(s.FirstName)
    s.LastName = p.Sanitize(s.LastName)
}