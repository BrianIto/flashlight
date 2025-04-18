package services

import (
	"flashlight/models"
	"gorm.io/gorm"
)

func CreateStudent(db *gorm.DB, name string, grade int) (models.Student, error) {
	student := models.Student{
		Name:  name,
		Grade: grade,
	}
	db.Create(&student)
	return student,  nil
}

func GetStudents(db *gorm.DB) ([]models.Student, error) {
	var students []models.Student
	err := db.Find(&students).Error
	return students, err
}

func UpdateStudent(db *gorm.DB, id uint, name string, grade int) error {
	var student models.Student
	if err := db.First(&student, id).Error; err != nil {
		return err
	}
	student.Name = name
	student.Grade = grade
	return db.Save(&student).Error
}

func DeleteStudent(db *gorm.DB, id uint) error {
	return db.Delete(&models.Student{}, id).Error
}

func GetStudentById (db *gorm.DB, id uint) (models.Student, error) {
	var user models.Student
	db.First(&user, "id = ?", id)
	return user, nil;
}
