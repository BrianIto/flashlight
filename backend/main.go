package main

import (
	"flashlight/controllers"
	"flashlight/database"
	_ "flashlight/docs"
	"flashlight/models"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Student API
// @version 1.0
// @description A simple CRUD API using Gin and GORM.
// @host localhost:8080
// @BasePath /

func main() {
	r := gin.Default()

	db := database.InitDatabase()
	db.AutoMigrate(&models.Student{})

	r.GET("/students", controllers.GetStudents)
	r.POST("/students", controllers.CreateStudent)
	r.PUT("/students", controllers.UpdateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/docs", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", []byte(`<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <title>Elements in HTML</title>
  
    <script src="https://unpkg.com/@stoplight/elements/web-components.min.js"></script>
    <link rel="stylesheet" href="https://unpkg.com/@stoplight/elements/styles.min.css">
  </head>
  <body>

    <elements-api
      apiDescriptionUrl="/swagger/doc.json"
      router="hash"
    />

  </body>
</html>`))
	})

	r.Run(":8080")

}
