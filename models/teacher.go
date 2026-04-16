package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name       string `json:"name" validate:"required"`
	Department string `json:"department"`
}
