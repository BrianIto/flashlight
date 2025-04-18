package controllers

import (
	"flashlight/database"
	"flashlight/models"
	"flashlight/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var db = database.InitDatabase()

// CreateStudent
// @Summary      Create a new student
// @Description  Adds a student to the database
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        student body models.StudentCreate true "Student data"
// @Success      201  {object}  models.Student
// @Router       /students [post]
func CreateStudent(c *gin.Context) {

	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	std, err := services.CreateStudent(db, student.Name, student.Grade)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, std)
}

// CreateStudent
// @Summary      Gets all students
// @Description  Get all students from the database
// @Tags         students
// @Produce      json
// @Success      200  {array}  models.Student
// @Router       /students [get]
func GetStudents(c *gin.Context) {
	students, err := services.GetStudents(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, students)
}

// UpdateStudent
// @Summary      Updates a student
// @Description  Edits a student from the database
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        student body models.StudentEdit true "Student data"
// @Success      201  {object}  models.Student
// @Router       /students [put]
func UpdateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err := services.UpdateStudent(db, student.ID, student.Name, student.Grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	newStd, err := services.GetStudentById(db, student.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, newStd)
}

// DeleteStudent
// @Summary      Deletes a student
// @Description  Deletes a student from the database
// @Tags         students
// @Param id path uint true "User ID"
// @Accept       json
// @Produce      json
// @Success      204  {object}  models.Student
// @Router       /students/{id} [delete]
func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID type"})
	}
	err1 := services.DeleteStudent(db, uint(idUint))
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Removed successfully"})

}
