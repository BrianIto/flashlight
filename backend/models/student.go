package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey"` // The ID of the Student
	Name  string `json:"name"`       // The name of the Student
	Grade int    `json:"grade"`      // The grade assigned to this student
}

type StudentCreate struct {
	Name  string `json:"name"`  // The ID of the Student
	Grade int    `json:"grade"` // The Grade of the Student
}

type StudentEdit struct {
	ID    uint   // The Id of the student
	Name  string `json:"name"`   // The Name of the student
	Grade int    `json: "grade"` // The Grade of the Student
}
