package models

import "gorm.io/gorm"

type Grade struct {
	gorm.Model
	Value     string  `json:"value"` // e.g., "A", "B+"
	StudentID uint    `json:"student_id"`
	Student   Student `gorm:"foreignKey:StudentID"` // Belongs to Student
	TeacherID uint    `json:"teacher_id"`
	Teacher   Teacher `gorm:"foreignKey:TeacherID"` // Belongs to Teacher
}
