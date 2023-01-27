package main

import (
	"api-sekolah/app/controller"
	"api-sekolah/app/repository"
	"api-sekolah/app/service"
	"api-sekolah/initialize"
	"api-sekolah/models"

	"github.com/gin-gonic/gin"
)

func init() {
	initialize.LoadDatabase()
}

func main() {
	router := gin.Default()

	studentRepository := repository.NewStudentRepository(initialize.DB)
	studentService := service.NewStudentService(studentRepository)
	studentController := controller.NewStudentController(studentService)

	v1 := router.Group("v1")
	v1.GET("/students", studentController.GetList)
	v1.POST("/students", studentController.StoreData)
	v1.GET("/students/:id", studentController.GetByID)
	v1.PUT("/students/:id", studentController.UpdateData)
	v1.DELETE("/students/:id", studentController.DeleteData)

	v1.GET("/test", func(c *gin.Context) {
		var students []models.Student
		// var class []models.Class

		initialize.DB.Debug().Find(&students)

		c.JSON(200, gin.H{
			"data": students,
		})
	})
	router.Run()
}
