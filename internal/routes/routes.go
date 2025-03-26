package routes

import (
	"github.com/gin-gonic/gin"
	"rest-project/internal/delivery"
)

func SetupRoutes(r *gin.Engine) {
	studentHandler := delivery.StudentHandler{}
	students := r.Group("api/v1/students")
	{
		students.GET("/", studentHandler.GetAllStudents)
		students.POST("/", studentHandler.CreateStudent)
		students.GET("/:id", studentHandler.GetStudent)
		students.PUT("/:id", studentHandler.UpdateStudent)
		students.DELETE("/:id", studentHandler.DeleteStudent)

	}

}
